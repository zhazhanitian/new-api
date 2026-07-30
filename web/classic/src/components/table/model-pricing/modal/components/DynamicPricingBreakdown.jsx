/*
Copyright (C) 2025 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/

import React from 'react';
import { Avatar, Tag, Table, Typography } from '@douyinfe/semi-ui';
import { IconPriceTag } from '@douyinfe/semi-icons';
import { parseTiersFromExpr, getCurrencyConfig } from '../../../../../helpers';
import { BILLING_PRICING_VARS } from '../../../../../constants';
import {
  splitBillingExprAndRequestRules,
  tryParseRequestRuleExpr,
  SOURCE_TIME,
  MATCH_RANGE,
  MATCH_EQ,
  MATCH_GTE,
  MATCH_LT,
  MATCH_CONTAINS,
  MATCH_EXISTS,
} from '../../../../../pages/Setting/Ratio/components/requestRuleExpr';

const { Text } = Typography;

const VAR_LABELS = { p: '输入', c: '输出' };

// 已知 param 附加费用项的中文标签（可按需扩展）
const PARAM_ADDON_LABELS = {
  image: '输入图片',
};

/**
 * 按顶层加法（深度0处的 +）拆分表达式，忽略括号内的加号。
 * 兼容 ' + '、' +\n'、'\n+ ' 等各种空白格式。
 */
function splitTopLevelPlus(expr) {
  const parts = [];
  let start = 0;
  let depth = 0;
  for (let i = 0; i < expr.length; i++) {
    const c = expr[i];
    if (c === '(') depth++;
    else if (c === ')') depth--;
    else if (depth === 0 && c === '+') {
      // 要求 + 前面是空白字符（避免误匹配科学计数法 1e+6）
      const prev = i > 0 ? expr[i - 1] : '';
      if (/\s/.test(prev)) {
        parts.push(expr.slice(start, i).trim());
        start = i + 1;
      }
    }
  }
  parts.push(expr.slice(start).trim());
  return parts.filter(Boolean);
}

/**
 * 从计费表达式中解析出 param(...) 条件附加费用项，例如：
 *   (param("image") != nil ? (param("image.#") > 0 ? param("image.#") : 1) * 13699 : 0)
 * 返回数组，每项形如 { paramName, label, unitCost }。
 */
function parseParamAddonCosts(exprStr) {
  if (!exprStr) return [];
  const addons = [];
  const segments = splitTopLevelPlus(exprStr);
  for (const seg of segments) {
    // 排除包含 tier( 的主档位项
    if (/tier\s*\(/.test(seg)) continue;
    // 匹配: param("X") != nil ? ... * COST : 0
    const m = seg.match(/param\("([^"]+)"\)\s*!=\s*nil[\s\S]*?\*\s*(\d+)[\s\S]*?:\s*0/);
    if (m) {
      const paramName = m[1];
      const unitCost = parseInt(m[2], 10);
      const label = PARAM_ADDON_LABELS[paramName] || `param(${paramName})`;
      addons.push({ paramName, label, unitCost });
    }
  }
  return addons;
}
const OP_LABELS = { '<': '<', '<=': '≤', '>': '>', '>=': '≥' };
const TIME_FUNC_LABELS = { hour: '小时', minute: '分钟', weekday: '星期', month: '月份', day: '日期' };

function formatTokenHint(value) {
  const n = Number(value);
  if (!Number.isFinite(n) || n === 0) return '';
  if (n >= 1000000) return `${(n / 1000000).toFixed(n % 1000000 === 0 ? 0 : 1)}M`;
  if (n >= 1000) return `${(n / 1000).toFixed(n % 1000 === 0 ? 0 : 1)}K`;
  return String(n);
}

function formatConditionSummary(conditions, t) {
  return conditions
    .map((c) => {
      if (c.var && c.op) {
        const varLabel = t(VAR_LABELS[c.var] || c.var);
        const hint = formatTokenHint(c.value);
        return `${varLabel} ${OP_LABELS[c.op] || c.op} ${hint || c.value}`;
      }
      return '';
    })
    .filter(Boolean)
    .join(' && ');
}


function describeCondition(cond, t) {
  if (cond.source === SOURCE_TIME) {
    const fn = t(TIME_FUNC_LABELS[cond.timeFunc] || cond.timeFunc);
    const tz = cond.timezone || 'UTC';
    if (cond.mode === MATCH_RANGE) {
      return `${fn} ${cond.rangeStart}:00~${cond.rangeEnd}:00 (${tz})`;
    }
    const opMap = { [MATCH_EQ]: '=', [MATCH_GTE]: '≥', [MATCH_LT]: '<' };
    return `${fn} ${opMap[cond.mode] || '='} ${cond.value} (${tz})`;
  }
  const src = cond.source === 'header' ? t('请求头') : t('请求参数');
  const path = cond.path || '';
  if (cond.mode === MATCH_EXISTS) return `${src} ${path} ${t('存在')}`;
  if (cond.mode === MATCH_CONTAINS) return `${src} ${path} ${t('包含')} "${cond.value}"`;
  const opMap = { eq: '=', gt: '>', gte: '≥', lt: '<', lte: '≤' };
  return `${src} ${path} ${opMap[cond.mode] || '='} ${cond.value}`;
}

function describeGroup(group, t) {
  const parts = (group.conditions || []).map((c) => describeCondition(c, t));
  return parts.join(' && ');
}

export default function DynamicPricingBreakdown({ billingExpr, t }) {
  const { symbol, rate } = getCurrencyConfig();
  const { billingExpr: baseExpr, requestRuleExpr: ruleExpr } =
    splitBillingExprAndRequestRules(billingExpr || '');

  const tiers = parseTiersFromExpr(baseExpr);
  const ruleGroups = tryParseRequestRuleExpr(ruleExpr || '');
  const addonCosts = parseParamAddonCosts(baseExpr);

  const hasTiers = tiers && tiers.length > 0;
  const hasRules = ruleGroups && ruleGroups.length > 0;
  const hasAddons = addonCosts && addonCosts.length > 0;

  if (!hasTiers && !hasRules && !hasAddons) {
    return (
      <div>
        <div className='flex items-center mb-3'>
          <Avatar size='small' color='amber' className='mr-2 shadow-md'>
            <IconPriceTag size={16} />
          </Avatar>
          <Text className='text-lg font-medium'>{t('动态计费')}</Text>
        </div>
        <div className='text-sm text-gray-500'>
          <code style={{ fontSize: 12, wordBreak: 'break-all' }}>{billingExpr}</code>
        </div>
      </div>
    );
  }

  const QUOTA_PER_USD = 500_000;
  // 和后端 QuotaRound 保持一致：先把表达式原始值换算成整数 quota，再转回货币
  const quotaToPrice = (rawCost) => (Math.round((rawCost / 1_000_000) * QUOTA_PER_USD) / QUOTA_PER_USD) * rate;
  const priceFields = BILLING_PRICING_VARS.map((v) => [v.field, v.shortLabel]);
  const hasFixedCost = hasTiers && tiers.some((tier) => tier.fixedCost > 0);

  const tierColumns = [
    {
      title: t('档位'),
      dataIndex: 'label',
      render: (text, record) => (
        <div>
          <Tag color='blue' size='small'>{text || t('默认')}</Tag>
          {record.condSummary && (
            <div className='text-xs text-gray-500 mt-1'>{record.condSummary}</div>
          )}
        </div>
      ),
    },
    ...priceFields
      .filter(([field]) => hasTiers && tiers.some((tier) => tier[field] > 0))
      .map(([field, label]) => ({
        title: `${t(label)} (${symbol}/1M tokens)`,
        dataIndex: field,
        render: (v) => v > 0 ? <Text strong>{`${symbol}${(v * rate).toFixed(4)}`}</Text> : '-',
      })),
    ...(hasFixedCost ? [{
      title: t('单次费用'),
      dataIndex: 'fixedCost',
      render: (v) => v > 0
        ? <Text strong>{`${symbol}${quotaToPrice(v).toFixed(6)}`}</Text>
        : '-',
    }] : []),
  ];

  const tierData = hasTiers
    ? tiers.map((tier, i) => ({
        key: `tier-${i}`,
        label: tier.label,
        condSummary: formatConditionSummary(tier.conditions, t),
        fixedCost: tier.fixedCost || 0,
        ...Object.fromEntries(priceFields.map(([field]) => [field, tier[field] || 0])),
      }))
    : [];

  return (
    <div>
      <div className='flex items-center mb-4'>
        <Avatar size='small' color='amber' className='mr-2 shadow-md'>
          <IconPriceTag size={16} />
        </Avatar>
        <div>
          <Text className='text-lg font-medium'>{t('动态计费')}</Text>
          <div className='text-xs text-gray-600'>
            {t('价格根据用量档位和请求条件动态调整')}
          </div>
        </div>
      </div>

      {hasTiers && (
        <div style={{ marginBottom: 16 }}>
          <Text strong className='text-sm' style={{ display: 'block', marginBottom: 8 }}>
            {t('分档价格表')}
          </Text>
          <Table
            dataSource={tierData}
            columns={tierColumns}
            pagination={false}
            size='small'
            bordered={false}
            className='!rounded-lg'
          />
        </div>
      )}

      {hasAddons && (
        <div style={{ marginBottom: 16 }}>
          <Text strong className='text-sm' style={{ display: 'block', marginBottom: 8 }}>
            {t('附加费用')}
          </Text>
          {addonCosts.map((addon) => (
            <div
              key={addon.paramName}
              style={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                padding: '8px 12px',
                borderRadius: 6,
                background: 'var(--semi-color-fill-0)',
                marginBottom: 4,
              }}
            >
              <Text size='small'>{t(addon.label)}</Text>
              <Text strong size='small'>
                {`${symbol}${quotaToPrice(addon.unitCost).toFixed(6)}/` + t('张')}
              </Text>
            </div>
          ))}
        </div>
      )}

      {hasRules && (
        <div style={{ marginBottom: 16 }}>
          <Text strong className='text-sm' style={{ display: 'block', marginBottom: 8 }}>
            {t('条件乘数')}
          </Text>
          {ruleGroups.map((group, gi) => (
            <div
              key={`group-${gi}`}
              style={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                padding: '8px 12px',
                borderRadius: 6,
                background: 'var(--semi-color-fill-0)',
                marginBottom: 4,
              }}
            >
              <Text size='small'>{describeGroup(group, t)}</Text>
              <Tag color='orange' size='small'>{group.multiplier}x</Tag>
            </div>
          ))}
        </div>
      )}

    </div>
  );
}
