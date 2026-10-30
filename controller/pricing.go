package controller

import (
	"encoding/json"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/channel/task/doubao"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/setting/ratio_setting"

	"github.com/gin-gonic/gin"
)

func attachSeedancePriceTiers(pricing []model.Pricing) {
	mappings := loadEnabledChannelModelMappings()
	for i := range pricing {
		tiers := seedanceDisplayTiers(pricing[i].ModelName, mappings)
		if len(tiers) == 0 {
			pricing[i].PriceTiers = nil
			continue
		}
		out := make([]model.PriceTier, 0, len(tiers))
		for _, t := range tiers {
			out = append(out, model.PriceTier{Label: t.Label, Ratio: t.Ratio})
		}
		pricing[i].PriceTiers = out
	}
}

// seedanceDisplayTiers 先按对外模型名查分档。查不到时，沿启用渠道的模型重定向找到官方 ID 再查。
// 广场分档表只内置了官方模型 ID，改名后的模型（如 doubao-seedance-2-0-fast-zj）否则会退回按量展示。
func seedanceDisplayTiers(modelName string, mappings []map[string]string) []doubao.PriceDisplayTier {
	if tiers := doubao.GetPriceDisplayTiers(modelName); len(tiers) > 0 {
		return tiers
	}
	upstream := mappedSeedanceModel(modelName, mappings)
	if upstream == "" || upstream == modelName {
		return nil
	}
	return doubao.GetPriceDisplayTiers(upstream)
}

func mappedSeedanceModel(modelName string, mappings []map[string]string) string {
	for _, mapping := range mappings {
		current := modelName
		seen := map[string]bool{current: true}
		for step := 0; step < 8; step++ {
			next := strings.TrimSpace(mapping[current])
			if next == "" || next == current || seen[next] {
				break
			}
			seen[next] = true
			current = next
			if len(doubao.GetPriceDisplayTiers(current)) > 0 {
				return current
			}
		}
	}
	return ""
}

func loadEnabledChannelModelMappings() []map[string]string {
	if model.DB == nil {
		return nil
	}
	var rows []struct {
		ModelMapping string
	}
	err := model.DB.Model(&model.Channel{}).
		Select("model_mapping").
		Where("status = ? AND model_mapping IS NOT NULL AND model_mapping != '' AND model_mapping != ?", common.ChannelStatusEnabled, "{}").
		Find(&rows).Error
	if err != nil {
		common.SysLog("load channel model mapping for price tiers failed: " + err.Error())
		return nil
	}
	mappings := make([]map[string]string, 0, len(rows))
	for _, row := range rows {
		parsed := map[string]string{}
		if json.Unmarshal([]byte(row.ModelMapping), &parsed) != nil || len(parsed) == 0 {
			continue
		}
		mappings = append(mappings, parsed)
	}
	return mappings
}

func filterPricingByUsableGroups(pricing []model.Pricing, usableGroup map[string]string) []model.Pricing {
	if len(pricing) == 0 {
		return pricing
	}
	if len(usableGroup) == 0 {
		return []model.Pricing{}
	}

	filtered := make([]model.Pricing, 0, len(pricing))
	for _, item := range pricing {
		if common.StringsContains(item.EnableGroup, "all") {
			filtered = append(filtered, item)
			continue
		}
		for _, group := range item.EnableGroup {
			if _, ok := usableGroup[group]; ok {
				filtered = append(filtered, item)
				break
			}
		}
	}
	return filtered
}

func GetPricing(c *gin.Context) {
	pricing := model.GetPricing()
	userId, exists := c.Get("id")
	usableGroup := map[string]string{}
	groupRatio := map[string]float64{}
	for s, f := range ratio_setting.GetGroupRatioCopy() {
		groupRatio[s] = f
	}
	var group string
	if exists {
		user, err := model.GetUserCache(userId.(int))
		if err == nil {
			group = user.Group
			for g := range groupRatio {
				ratio, ok := ratio_setting.GetGroupGroupRatio(group, g)
				if ok {
					groupRatio[g] = ratio
				}
			}
		}
	}

	usableGroup = service.GetUserUsableGroups(group)
	pricing = filterPricingByUsableGroups(pricing, usableGroup)
	attachSeedancePriceTiers(pricing)
	// check groupRatio contains usableGroup
	for group := range ratio_setting.GetGroupRatioCopy() {
		if _, ok := usableGroup[group]; !ok {
			delete(groupRatio, group)
		}
	}

	c.JSON(200, gin.H{
		"success":            true,
		"data":               pricing,
		"vendors":            model.GetVendors(),
		"group_ratio":        groupRatio,
		"usable_group":       usableGroup,
		"supported_endpoint": model.GetSupportedEndpointMap(),
		"auto_groups":        service.GetUserAutoGroup(group),
		"pricing_version":    "a42d372ccf0b5dd13ecf71203521f9d2",
	})
}

func ResetModelRatio(c *gin.Context) {
	defaultStr := ratio_setting.DefaultModelRatio2JSONString()
	err := model.UpdateOption("ModelRatio", defaultStr)
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	err = ratio_setting.UpdateModelRatioByJSONString(defaultStr)
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "重置模型倍率成功",
	})
}
