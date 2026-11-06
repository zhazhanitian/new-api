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
import { Typography, Card, Button, Input } from '@douyinfe/semi-ui';
import { Copy, Users, BarChart2, TrendingUp, Zap } from 'lucide-react';

const { Text } = Typography;

const InvitationCard = ({
  t,
  userState,
  renderQuota,
  setOpenTransfer,
  affLink,
  handleAffLinkClick,
  complianceConfirmed = true,
}) => {
  return (
    <Card className='!rounded-2xl shadow-sm border-0'>
      <div className='mb-5 flex items-center justify-between gap-3'>
        <Typography.Title heading={5} style={{ margin: 0 }}>
          {t('邀请奖励')}
        </Typography.Title>
        <Button
          type='primary'
          theme='light'
          size='small'
          disabled={
            !complianceConfirmed ||
            !userState?.user?.aff_quota ||
            userState?.user?.aff_quota <= 0
          }
          onClick={() => setOpenTransfer(true)}
          icon={<Zap size={14} />}
        >
          {t('划转到余额')}
        </Button>
      </div>

      <div className='flex flex-col gap-5'>
        <div
          className='grid grid-cols-3 overflow-hidden rounded-2xl'
          style={{ border: '1px solid var(--semi-color-border)' }}
        >
          {[
            {
              label: t('待使用收益'),
              value: renderQuota(userState?.user?.aff_quota || 0),
              icon: <TrendingUp size={14} />,
            },
            {
              label: t('总收益'),
              value: renderQuota(userState?.user?.aff_history_quota || 0),
              icon: <BarChart2 size={14} />,
            },
            {
              label: t('邀请人数'),
              value: userState?.user?.aff_count || 0,
              icon: <Users size={14} />,
            },
          ].map((item, index) => (
            <div
              key={item.label}
              className='px-3 py-4'
              style={{
                background: 'var(--semi-color-fill-0)',
                borderLeft:
                  index === 0 ? 'none' : '1px solid var(--semi-color-border)',
              }}
            >
              <div
                className='mb-2 flex items-center gap-1 text-xs'
                style={{ color: 'var(--semi-color-text-2)' }}
              >
                {item.icon}
                {item.label}
              </div>
              <div
                className='truncate text-lg font-semibold'
                style={{ color: 'var(--semi-color-text-0)' }}
              >
                {item.value}
              </div>
            </div>
          ))}
        </div>

        {!complianceConfirmed && (
          <Text type='tertiary' size='small'>
            {t('邀请奖励划转已禁用，管理员需先确认合规声明。')}
          </Text>
        )}

        <div>
          <Text strong className='mb-3 block'>
            {t('邀请链接')}
          </Text>
          <Input
            value={affLink}
            readonly
            suffix={
              <Button
                type='primary'
                theme='solid'
                onClick={handleAffLinkClick}
                icon={<Copy size={14} />}
              >
                {t('复制')}
              </Button>
            }
          />
        </div>

        <div
          className='rounded-2xl px-4 py-4'
          style={{ border: '1px solid var(--semi-color-border)' }}
        >
          <Text strong className='mb-3 block'>
            {t('奖励说明')}
          </Text>
          <div className='flex flex-col gap-2'>
            {[
              t('邀请好友注册，好友充值后您可获得相应奖励'),
              t('通过划转功能将奖励额度转入到您的账户余额中'),
              t('邀请的好友越多，获得的奖励越多'),
            ].map((line) => (
              <div key={line} className='flex items-start gap-2'>
                <span
                  className='mt-1.5 h-1.5 w-1.5 shrink-0 rounded-full'
                  style={{ background: 'var(--semi-color-success)' }}
                />
                <Text type='tertiary' size='small'>
                  {line}
                </Text>
              </div>
            ))}
          </div>
        </div>
      </div>
    </Card>
  );
};

export default InvitationCard;
