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
import React, { useState, useEffect, useMemo } from 'react';
import {
  Modal,
  Table,
  Typography,
  Toast,
  Empty,
  Button,
  Input,
  Tag,
} from '@douyinfe/semi-ui';
import {
  IllustrationNoResult,
  IllustrationNoResultDark,
} from '@douyinfe/semi-illustrations';
import { IconSearch } from '@douyinfe/semi-icons';
import { API, timestamp2string } from '../../../helpers';
import { isAdmin } from '../../../helpers/utils';
import { useIsMobile } from '../../../hooks/common/useIsMobile';
const { Text } = Typography;

// 状态映射配置
const STATUS_CONFIG = {
  success: { color: 'green', key: '成功' },
  pending: { color: 'orange', key: '待支付' },
  failed: { color: 'red', key: '失败' },
  expired: { color: 'red', key: '已过期' },
};

// 支付方式映射
const PAYMENT_METHOD_MAP = {
  stripe: 'Stripe',
  creem: 'Creem',
  waffo: 'Waffo',
  alipay: '支付宝',
  wxpay: '微信支付',
  paycenter: '微信支付',
};

const columnTitle = (text) => (
  <span style={{ whiteSpace: 'nowrap' }}>{text}</span>
);

const formatYuan = (value) => {
  const amount = Number(value);
  if (!Number.isFinite(amount)) return '-';
  return `¥${amount.toFixed(2)}`;
};

const TopupHistoryModal = ({ visible, onCancel, t }) => {
  const [loading, setLoading] = useState(false);
  const [topups, setTopups] = useState([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [keyword, setKeyword] = useState('');
  const isMobile = useIsMobile();

  const loadTopups = async (currentPage, currentPageSize) => {
    setLoading(true);
    try {
      const base = isAdmin() ? '/api/user/topup' : '/api/user/topup/self';
      const qs =
        `p=${currentPage}&page_size=${currentPageSize}` +
        (keyword ? `&keyword=${encodeURIComponent(keyword)}` : '');
      const endpoint = `${base}?${qs}`;
      const res = await API.get(endpoint);
      const { success, message, data } = res.data;
      if (success) {
        setTopups(data.items || []);
        setTotal(data.total || 0);
      } else {
        Toast.error({ content: message || t('加载失败') });
      }
    } catch (error) {
      Toast.error({ content: t('加载账单失败') });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (visible) {
      loadTopups(page, pageSize);
    }
  }, [visible, page, pageSize, keyword]);

  const handlePageChange = (currentPage) => {
    setPage(currentPage);
  };

  const handlePageSizeChange = (currentPageSize) => {
    setPageSize(currentPageSize);
    setPage(1);
  };

  const handleKeywordChange = (value) => {
    setKeyword(value);
    setPage(1);
  };

  // 管理员补单
  const handleAdminComplete = async (tradeNo) => {
    try {
      const res = await API.post('/api/user/topup/complete', {
        trade_no: tradeNo,
      });
      const { success, message } = res.data;
      if (success) {
        Toast.success({ content: t('补单成功') });
        await loadTopups(page, pageSize);
      } else {
        Toast.error({ content: message || t('补单失败') });
      }
    } catch (e) {
      Toast.error({ content: t('补单失败') });
    }
  };

  const confirmAdminComplete = (tradeNo) => {
    Modal.confirm({
      title: t('确认补单'),
      content: t('是否将该订单标记为成功并为用户入账？'),
      onOk: () => handleAdminComplete(tradeNo),
    });
  };

  // 渲染状态徽章
  const renderStatusBadge = (status) => {
    const config = STATUS_CONFIG[status] || { color: 'grey', key: status };
    return (
      <Tag color={config.color} size='small'>
        {t(config.key)}
      </Tag>
    );
  };

  // 渲染支付方式
  const renderPaymentMethod = (pm) => {
    const displayName = PAYMENT_METHOD_MAP[pm];
    return <Text>{displayName ? t(displayName) : pm || '-'}</Text>;
  };

  const isSubscriptionTopup = (record) => {
    const tradeNo = (record?.trade_no || '').toLowerCase();
    return Number(record?.amount || 0) === 0 && tradeNo.startsWith('sub');
  };

  // 检查是否为管理员
  const userIsAdmin = useMemo(() => isAdmin(), []);

  const columns = useMemo(() => {
    const baseColumns = [
      {
        title: columnTitle(t('时间')),
        dataIndex: 'create_time',
        key: 'create_time',
        width: 176,
        render: (time) => (
          <Text style={{ whiteSpace: 'nowrap' }}>{timestamp2string(time)}</Text>
        ),
      },
      ...(userIsAdmin
        ? [
            {
              title: columnTitle(t('用户')),
              dataIndex: 'user_id',
              key: 'user_id',
              width: 72,
              render: (userId) => <Text>{userId ?? '-'}</Text>,
            },
          ]
        : []),
      {
        title: columnTitle(t('支付方式')),
        dataIndex: 'payment_method',
        key: 'payment_method',
        width: 110,
        render: renderPaymentMethod,
      },
      {
        title: columnTitle(t('金额')),
        dataIndex: 'money',
        key: 'money',
        width: 120,
        render: (money, record) => {
          if (isSubscriptionTopup(record)) {
            return (
              <Tag color='purple' size='small'>
                {t('订阅套餐')}
              </Tag>
            );
          }
          const paid = Number(money || 0);
          const quota = Number(record?.amount || 0);
          const showQuota = quota > 0 && Math.abs(quota - paid) > 0.001;
          return (
            <div>
              <Text strong>{formatYuan(paid)}</Text>
              {showQuota && (
                <div
                  className='text-xs'
                  style={{ color: 'var(--semi-color-text-2)' }}
                >
                  {t('额度')} {formatYuan(quota)}
                </div>
              )}
            </div>
          );
        },
      },
      {
        title: columnTitle(t('状态')),
        dataIndex: 'status',
        key: 'status',
        width: 88,
        render: renderStatusBadge,
      },
      {
        title: columnTitle(t('订单号')),
        dataIndex: 'trade_no',
        key: 'trade_no',
        render: (text) => (
          <Text ellipsis={{ showTooltip: true }} copyable>
            {text}
          </Text>
        ),
      },
    ];

    const hasPending = topups.some((record) => record.status === 'pending');
    if (userIsAdmin && hasPending) {
      baseColumns.push({
        title: columnTitle(t('操作')),
        key: 'action',
        width: 88,
        render: (_, record) => {
          if (record.status !== 'pending') return null;
          return (
            <Button
              size='small'
              type='primary'
              theme='light'
              onClick={() => confirmAdminComplete(record.trade_no)}
            >
              {t('补单')}
            </Button>
          );
        },
      });
    }

    return baseColumns;
  }, [t, userIsAdmin, topups]);

  return (
    <Modal
      title={t('充值账单')}
      visible={visible}
      onCancel={onCancel}
      footer={null}
      width={isMobile ? '100%' : 920}
      bodyStyle={{ paddingBottom: 16 }}
    >
      <div className='mb-3'>
        <Input
          prefix={<IconSearch />}
          placeholder={t('订单号')}
          value={keyword}
          onChange={handleKeywordChange}
          showClear
        />
      </div>
      <Table
        columns={columns}
        dataSource={topups}
        loading={loading}
        rowKey='id'
        pagination={{
          currentPage: page,
          pageSize: pageSize,
          total: total,
          showSizeChanger: true,
          pageSizeOpts: [10, 20, 50, 100],
          onPageChange: handlePageChange,
          onPageSizeChange: handlePageSizeChange,
        }}
        empty={
          <Empty
            image={<IllustrationNoResult style={{ width: 150, height: 150 }} />}
            darkModeImage={
              <IllustrationNoResultDark style={{ width: 150, height: 150 }} />
            }
            description={t('暂无充值记录')}
            style={{ padding: 30 }}
          />
        }
      />
    </Modal>
  );
};

export default TopupHistoryModal;
