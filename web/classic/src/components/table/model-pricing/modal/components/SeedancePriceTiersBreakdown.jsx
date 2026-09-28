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

import React, { useMemo } from 'react';
import { Avatar, Tag, Table, Typography } from '@douyinfe/semi-ui';
import { IconPriceTag } from '@douyinfe/semi-icons';
import { calculateModelPrice } from '../../../../../helpers';

const { Text } = Typography;

function splitTierLabel(label) {
  const text = String(label || '').trim();
  if (!text) return { title: '', detail: '' };
  const parts = text.split('·').map((s) => s.trim()).filter(Boolean);
  if (parts.length >= 2) {
    return { title: parts[0], detail: parts.slice(1).join(' · ') };
  }
  return { title: text, detail: '' };
}

/**
 * Seedance 等视频模型的分档价展示（样式对齐动态计费「分档价格表」）。
 * 仅展示用：价格 = 后台基准输入价 × price_tiers.ratio × 分组倍率。
 */
export default function SeedancePriceTiersBreakdown({
  modelData,
  groupRatio,
  currency,
  siteDisplayType,
  tokenUnit,
  displayPrice,
  selectedGroup,
  t,
}) {
  const tiers = Array.isArray(modelData?.price_tiers)
    ? modelData.price_tiers
    : [];

  const priceData = useMemo(() => {
    if (!modelData || tiers.length === 0) return null;
    return calculateModelPrice({
      record: modelData,
      selectedGroup: selectedGroup || 'all',
      groupRatio,
      tokenUnit,
      displayPrice,
      currency,
      quotaDisplayType: siteDisplayType,
    });
  }, [
    modelData,
    tiers.length,
    selectedGroup,
    groupRatio,
    tokenUnit,
    displayPrice,
    currency,
    siteDisplayType,
  ]);

  if (!priceData?.priceTiers?.length) {
    return null;
  }

  const unitLabel = priceData.unitLabel || 'M';
  const unitTitle =
    siteDisplayType === 'TOKENS'
      ? t('倍率')
      : `${t('价格')} (/1${unitLabel} Tokens)`;

  const columns = [
    {
      title: t('档位'),
      dataIndex: 'title',
      render: (text, record) => (
        <div>
          <Tag color='blue' size='small'>
            {text || t('默认')}
          </Tag>
          {record.detail ? (
            <div className='text-xs text-gray-500 mt-1'>{record.detail}</div>
          ) : null}
        </div>
      ),
    },
    {
      title: unitTitle,
      dataIndex: 'value',
      align: 'right',
      render: (v) => (
        <Text strong className='text-orange-600'>
          {v}
        </Text>
      ),
    },
  ];

  const dataSource = priceData.priceTiers.map((tier) => {
    const { title, detail } = splitTierLabel(tier.label);
    return {
      key: tier.key,
      title,
      detail,
      value: tier.value,
    };
  });

  return (
    <div>
      <div className='flex items-center mb-4'>
        <Avatar size='small' color='amber' className='mr-2 shadow-md'>
          <IconPriceTag size={16} />
        </Avatar>
        <div>
          <Text className='text-lg font-medium'>{t('分档计费')}</Text>
          <div className='text-xs text-gray-600'>
            {t('价格随是否含视频输入、分辨率等条件变化')}
          </div>
        </div>
      </div>

      <Text strong className='text-sm' style={{ display: 'block', marginBottom: 8 }}>
        {t('分档价格表')}
      </Text>
      <Table
        dataSource={dataSource}
        columns={columns}
        pagination={false}
        size='small'
        bordered={false}
        className='!rounded-lg'
      />
    </div>
  );
}
