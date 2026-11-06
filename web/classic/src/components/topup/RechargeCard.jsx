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

import React, { useEffect, useRef, useState } from 'react';
import {
  Typography,
  Card,
  Button,
  Banner,
  Skeleton,
  Form,
  Spin,
  Tooltip,
  Tag,
  Tabs,
  TabPane,
} from '@douyinfe/semi-ui';
import { SiAlipay, SiWechat, SiStripe } from 'react-icons/si';
import {
  CreditCard,
  Wallet,
  BarChart2,
  TrendingUp,
  Receipt,
  Sparkles,
} from 'lucide-react';
import { IconGift } from '@douyinfe/semi-icons';
import { useMinimumLoadingTime } from '../../hooks/common/useMinimumLoadingTime';
import { getCurrencyConfig } from '../../helpers/render';
import SubscriptionPlansCard from './SubscriptionPlansCard';

const { Text } = Typography;

const RechargeCard = ({
  t,
  enableOnlineTopUp,
  enableStripeTopUp,
  enableCreemTopUp,
  creemProducts,
  creemPreTopUp,
  presetAmounts,
  selectedPreset,
  selectPresetAmount,
  formatLargeNumber,
  priceRatio,
  topUpCount,
  minTopUp,
  renderQuotaWithAmount,
  getAmount,
  quoteAmount,
  setTopUpCount,
  setSelectedPreset,
  renderAmount,
  amountLoading,
  payMethods,
  preTopUp,
  paymentLoading,
  payWay,
  redemptionCode,
  setRedemptionCode,
  topUp,
  isSubmitting,
  topUpLink,
  openTopUpLink,
  userState,
  renderQuota,
  statusLoading,
  topupInfo,
  onOpenHistory,
  enableWaffoTopUp,
  enableWaffoPancakeTopUp,
  enablePayCenterTopUp,
  subscriptionLoading = false,
  subscriptionPlans = [],
  billingPreference,
  onChangeBillingPreference,
  activeSubscriptions = [],
  allSubscriptions = [],
  reloadSubscriptionSelf,
  enableRedemption = true,
}) => {
  const onlineFormApiRef = useRef(null);
  const redeemFormApiRef = useRef(null);
  const initialTabSetRef = useRef(false);
  const showAmountSkeleton = useMinimumLoadingTime(amountLoading);
  const [activeTab, setActiveTab] = useState('topup');
  const shouldShowSubscription =
    !subscriptionLoading && subscriptionPlans.length > 0;
  const [pickedPayMethod, setPickedPayMethod] = useState('');

  const isPayMethodDisabled = (payMethod) => {
    const minTopupVal = Number(payMethod.min_topup) || 0;
    const isStripe = payMethod.type === 'stripe';
    const isWaffo =
      typeof payMethod.type === 'string' && payMethod.type.startsWith('waffo:');
    const isWaffoPancake = payMethod.type === 'waffo_pancake';
    const isPayCenter = payMethod.type === 'paycenter';
    return (
      (!enableOnlineTopUp &&
        !isStripe &&
        !isWaffo &&
        !isWaffoPancake &&
        !isPayCenter) ||
      (!enableStripeTopUp && isStripe) ||
      (!enableWaffoTopUp && isWaffo) ||
      (!enableWaffoPancakeTopUp && isWaffoPancake) ||
      (!enablePayCenterTopUp && isPayCenter) ||
      minTopupVal > Number(topUpCount || 0)
    );
  };

  const regularPayMethods = (payMethods || []).filter((method) => {
    if (enablePayCenterTopUp && method.type === 'wxpay') return false;
    return !isPayMethodDisabled(method);
  });

  const activePayMethod = regularPayMethods.some(
    (method) =>
      method.type === pickedPayMethod && !isPayMethodDisabled(method),
  )
    ? pickedPayMethod
    : regularPayMethods.find((method) => !isPayMethodDisabled(method))?.type ||
      '';

  const showYuanFace =
    enablePayCenterTopUp &&
    !enableOnlineTopUp &&
    !enableStripeTopUp &&
    !enableWaffoTopUp &&
    !enableWaffoPancakeTopUp &&
    getCurrencyConfig().type === 'CNY';

  const refreshQuote = (value, method) => {
    const payMethod = method || activePayMethod;
    if (quoteAmount && payMethod) {
      return quoteAmount(payMethod, value);
    }
    return getAmount(value);
  };

  const renderPayMethodIcon = (payMethod) => {
    if (payMethod.type === 'alipay') {
      return <SiAlipay size={18} color='#1677FF' />;
    }
    if (payMethod.type === 'wxpay' || payMethod.type === 'paycenter') {
      return <SiWechat size={18} color='#07C160' />;
    }
    if (payMethod.type === 'stripe') {
      return <SiStripe size={18} color='#635BFF' />;
    }
    if (payMethod.icon) {
      return (
        <img
          src={payMethod.icon}
          alt=''
          style={{ width: 18, height: 18, objectFit: 'contain' }}
        />
      );
    }
    if (payMethod.type === 'waffo_pancake') {
      return <CreditCard size={18} color='var(--semi-color-primary)' />;
    }
    return (
      <CreditCard
        size={18}
        color={payMethod.color || 'var(--semi-color-text-2)'}
      />
    );
  };

  useEffect(() => {
    if (initialTabSetRef.current) return;
    if (subscriptionLoading) return;
    setActiveTab(shouldShowSubscription ? 'subscription' : 'topup');
    initialTabSetRef.current = true;
  }, [shouldShowSubscription, subscriptionLoading]);

  useEffect(() => {
    if (!shouldShowSubscription && activeTab !== 'topup') {
      setActiveTab('topup');
    }
  }, [shouldShowSubscription, activeTab]);
  const topupContent = (
    <div className='flex flex-col gap-5'>
      <div
        className='grid grid-cols-3 overflow-hidden rounded-2xl'
        style={{ border: '1px solid var(--semi-color-border)' }}
      >
        {[
          {
            label: t('当前余额'),
            value: renderQuota(userState?.user?.quota),
            icon: <Wallet size={14} />,
          },
          {
            label: t('历史消耗'),
            value: renderQuota(userState?.user?.used_quota),
            icon: <TrendingUp size={14} />,
          },
          {
            label: t('请求次数'),
            value: userState?.user?.request_count || 0,
            icon: <BarChart2 size={14} />,
          },
        ].map((item, index) => (
          <div
            key={item.label}
            className='px-3 py-4 sm:px-4'
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
              className='truncate text-lg font-semibold sm:text-2xl'
              style={{ color: 'var(--semi-color-text-0)' }}
            >
              {item.value}
            </div>
          </div>
        ))}
      </div>

      {statusLoading ? (
        <div className='flex justify-center py-8'>
          <Spin size='large' />
        </div>
      ) : enableOnlineTopUp ||
        enableStripeTopUp ||
        enableCreemTopUp ||
        enableWaffoTopUp ||
        enableWaffoPancakeTopUp ||
        enablePayCenterTopUp ? (
        <Form
          getFormApi={(api) => (onlineFormApiRef.current = api)}
          initValues={{ topUpCount: topUpCount }}
        >
          <div className='flex flex-col gap-5'>
            {(enableOnlineTopUp ||
              enableStripeTopUp ||
              enableWaffoTopUp ||
              enablePayCenterTopUp) && (
              <div>
                <div className='mb-3 flex flex-wrap items-center gap-2'>
                  <Text strong>{t('选择额度')}</Text>
                  {!showYuanFace &&
                    (() => {
                      const { symbol, rate, type } = getCurrencyConfig();
                      if (type === 'USD') return null;
                      return (
                        <Text type='tertiary' size='small'>
                          1 $ = {rate.toFixed(2)} {symbol}
                        </Text>
                      );
                    })()}
                </div>
                <div className='grid grid-cols-2 gap-3 sm:grid-cols-3'>
                  {presetAmounts.map((preset, index) => {
                    const discount =
                      preset.discount ||
                      topupInfo?.discount?.[preset.value] ||
                      1.0;
                    const originalPrice = preset.value * priceRatio;
                    const discountedPrice = originalPrice * discount;
                    const hasDiscount = discount < 1.0;
                    const actualPay = discountedPrice;

                    const { symbol, rate, type } = getCurrencyConfig();
                    const statusStr = localStorage.getItem('status');
                    let usdRate = 7;
                    try {
                      if (statusStr) {
                        const s = JSON.parse(statusStr);
                        usdRate = s?.usd_exchange_rate || 7;
                      }
                    } catch (e) {}

                    let displayValue = preset.value;
                    let displayActualPay = actualPay;

                    if (showYuanFace) {
                      displayValue = preset.value;
                      displayActualPay = preset.value * discount;
                    } else if (type === 'USD') {
                      displayActualPay = actualPay / usdRate;
                    } else if (type === 'CNY') {
                      displayValue = preset.value * usdRate;
                    } else if (type === 'CUSTOM') {
                      displayValue = preset.value * rate;
                      displayActualPay = (actualPay / usdRate) * rate;
                    }

                    const selected = selectedPreset === preset.value;
                    return (
                      <button
                        key={index}
                        type='button'
                        onClick={() => {
                          selectPresetAmount(preset);
                          onlineFormApiRef.current?.setValue(
                            'topUpCount',
                            preset.value,
                          );
                          refreshQuote(preset.value);
                        }}
                        className='rounded-xl px-3 py-4 text-left'
                        style={{
                          border: selected
                            ? '2px solid var(--semi-color-primary)'
                            : '1px solid var(--semi-color-border)',
                          background: selected
                            ? 'var(--semi-color-primary-light-default)'
                            : 'var(--semi-color-bg-0)',
                        }}
                      >
                        <div className='flex items-center justify-between gap-2'>
                          <span
                            className='text-lg font-semibold'
                            style={{ color: 'var(--semi-color-text-0)' }}
                          >
                            {symbol}
                            {formatLargeNumber(displayValue)}
                          </span>
                          {hasDiscount && (
                            <Tag color='green' size='small'>
                              {t('折').includes('off')
                                ? ((1 - parseFloat(discount)) * 100).toFixed(1)
                                : (discount * 10).toFixed(1)}
                              {t('折')}
                            </Tag>
                          )}
                        </div>
                        {hasDiscount && (
                          <div
                            className='mt-1 text-xs'
                            style={{ color: 'var(--semi-color-text-2)' }}
                          >
                            {t('实付')} {symbol}
                            {displayActualPay.toFixed(2)}
                          </div>
                        )}
                      </button>
                    );
                  })}
                </div>
              </div>
            )}

            {(enableOnlineTopUp ||
              enableStripeTopUp ||
              enableWaffoTopUp ||
              enableWaffoPancakeTopUp ||
              enablePayCenterTopUp) && (
              <Form.InputNumber
                field='topUpCount'
                label={t(showYuanFace ? '充值金额' : '自定义数量')}
                disabled={
                  !enableOnlineTopUp &&
                  !enableStripeTopUp &&
                  !enableWaffoTopUp &&
                  !enableWaffoPancakeTopUp &&
                  !enablePayCenterTopUp
                }
                placeholder={t('输入数量')}
                value={topUpCount}
                min={minTopUp}
                max={999999999}
                step={1}
                precision={0}
                onChange={async (value) => {
                  if (value && value >= 1) {
                    setTopUpCount(value);
                    setSelectedPreset(null);
                    await refreshQuote(value);
                  }
                }}
                onBlur={(e) => {
                  const value = parseInt(e.target.value);
                  if (!value || value < 1) {
                    setTopUpCount(1);
                    refreshQuote(1);
                  }
                }}
                formatter={(value) => (value ? `${value}` : '')}
                parser={(value) =>
                  value ? parseInt(value.replace(/[^\d]/g, '')) : 0
                }
                extraText={
                  <Text type='tertiary' size='small'>
                    {t('最低')} {renderQuotaWithAmount(minTopUp)}
                  </Text>
                }
                style={{ width: '100%' }}
              />
            )}

            {regularPayMethods.length > 0 && (
              <div>
                <div className='mb-3'>
                  <Text strong>{t('支付方式')}</Text>
                </div>
                <div className='grid grid-cols-1 gap-3 sm:grid-cols-2'>
                  {regularPayMethods.map((payMethod) => {
                    const disabled = isPayMethodDisabled(payMethod);
                    const selected = activePayMethod === payMethod.type;
                    const minTopupVal = Number(payMethod.min_topup) || 0;
                    const buttonEl = (
                      <button
                        type='button'
                        disabled={disabled}
                        onClick={() => {
                          if (disabled) return;
                          setPickedPayMethod(payMethod.type);
                          refreshQuote(topUpCount, payMethod.type);
                        }}
                        className='flex w-full items-center gap-3 rounded-xl px-3 py-3 text-left'
                        style={{
                          border: selected
                            ? '2px solid var(--semi-color-primary)'
                            : '1px solid var(--semi-color-border)',
                          background: selected
                            ? 'var(--semi-color-primary-light-default)'
                            : 'var(--semi-color-bg-0)',
                          opacity: disabled ? 0.45 : 1,
                          cursor: disabled ? 'not-allowed' : 'pointer',
                        }}
                      >
                        {renderPayMethodIcon(payMethod)}
                        <span
                          className='min-w-0 flex-1 truncate text-sm font-medium'
                          style={{ color: 'var(--semi-color-text-0)' }}
                        >
                          {payMethod.name}
                        </span>
                        <span
                          className='flex h-4 w-4 shrink-0 items-center justify-center rounded-full'
                          style={{
                            border: selected
                              ? 'none'
                              : '1px solid var(--semi-color-border)',
                            background: selected
                              ? 'var(--semi-color-primary)'
                              : 'transparent',
                          }}
                        >
                          {selected && (
                            <span className='h-1.5 w-1.5 rounded-full bg-white' />
                          )}
                        </span>
                      </button>
                    );

                    if (disabled && minTopupVal > Number(topUpCount || 0)) {
                      return (
                        <Tooltip
                          content={
                            t('此支付方式最低充值金额为') + ' ' + minTopupVal
                          }
                          key={payMethod.type}
                        >
                          <span className='block'>{buttonEl}</span>
                        </Tooltip>
                      );
                    }
                    return (
                      <React.Fragment key={payMethod.type}>
                        {buttonEl}
                      </React.Fragment>
                    );
                  })}
                </div>
              </div>
            )}

            {enableCreemTopUp && creemProducts.length > 0 && (
              <div>
                <div className='mb-3'>
                  <Text strong>{t('Creem 充值')}</Text>
                </div>
                <div className='grid grid-cols-1 gap-3 sm:grid-cols-2'>
                  {creemProducts.map((product, index) => (
                    <button
                      key={index}
                      type='button'
                      onClick={() => creemPreTopUp(product)}
                      className='rounded-xl px-4 py-4 text-left'
                      style={{
                        border: '1px solid var(--semi-color-border)',
                        background: 'var(--semi-color-bg-0)',
                      }}
                    >
                      <div
                        className='font-medium'
                        style={{ color: 'var(--semi-color-text-0)' }}
                      >
                        {product.name}
                      </div>
                      <div
                        className='mt-1 text-sm'
                        style={{ color: 'var(--semi-color-text-2)' }}
                      >
                        {t('充值额度')}: {product.quota}
                      </div>
                      <div
                        className='mt-2 text-lg font-semibold'
                        style={{ color: 'var(--semi-color-primary)' }}
                      >
                        {product.currency === 'EUR' ? '€' : '$'}
                        {product.price}
                      </div>
                    </button>
                  ))}
                </div>
              </div>
            )}

            {regularPayMethods.length > 0 && (
              <div
                className='flex flex-col gap-4 rounded-2xl px-4 py-4 sm:flex-row sm:items-center sm:justify-between'
                style={{ background: 'var(--semi-color-fill-0)' }}
              >
                <div>
                  <div
                    className='mb-1 text-xs'
                    style={{ color: 'var(--semi-color-text-2)' }}
                  >
                    {t('应付金额')}
                  </div>
                  <Skeleton
                    loading={showAmountSkeleton}
                    active
                    placeholder={
                      <Skeleton.Title
                        style={{ width: 120, height: 28, borderRadius: 6 }}
                      />
                    }
                  >
                    <div
                      className='text-2xl font-semibold'
                      style={{ color: 'var(--semi-color-text-0)' }}
                    >
                      {renderAmount()}
                    </div>
                  </Skeleton>
                </div>
                <Button
                  theme='solid'
                  type='primary'
                  size='large'
                  htmlType='button'
                  disabled={!activePayMethod}
                  loading={paymentLoading && payWay === activePayMethod}
                  onClick={() => preTopUp(activePayMethod)}
                  className='!rounded-xl sm:!px-8'
                >
                  {t('立即充值')}
                </Button>
              </div>
            )}
          </div>
        </Form>
      ) : (
        <Banner
          type='info'
          description={t(
            '管理员未开启在线充值功能，请联系管理员开启或使用兑换码充值。',
          )}
          className='!rounded-xl'
          closeIcon={null}
        />
      )}

      {enableRedemption ? (
        <div
          className='rounded-2xl px-4 py-4'
          style={{ border: '1px solid var(--semi-color-border)' }}
        >
          <Text strong className='mb-3 block'>
            {t('兑换码')}
          </Text>
          <Form
            getFormApi={(api) => (redeemFormApiRef.current = api)}
            initValues={{ redemptionCode: redemptionCode }}
          >
            <Form.Input
              field='redemptionCode'
              noLabel={true}
              placeholder={t('请输入兑换码')}
              value={redemptionCode}
              onChange={(value) => setRedemptionCode(value)}
              prefix={<IconGift />}
              suffix={
                <Button
                  type='primary'
                  theme='solid'
                  htmlType='button'
                  onClick={topUp}
                  loading={isSubmitting}
                >
                  {t('兑换')}
                </Button>
              }
              showClear
              style={{ width: '100%' }}
              extraText={
                topUpLink && (
                  <Text type='tertiary'>
                    {t('在找兑换码？')}
                    <Text
                      type='secondary'
                      underline
                      className='cursor-pointer'
                      onClick={openTopUpLink}
                    >
                      {t('购买兑换码')}
                    </Text>
                  </Text>
                )
              }
            />
          </Form>
        </div>
      ) : (
        <Banner
          type='warning'
          description={t('兑换码功能已禁用，管理员需先确认合规声明。')}
          closeIcon={null}
          className='!rounded-xl'
        />
      )}
    </div>
  );


  return (
    <Card className='!rounded-2xl shadow-sm border-0'>
      {/* 卡片头部 */}
      <div className='mb-5 flex items-center justify-between'>
        <Typography.Title heading={5} style={{ margin: 0 }}>
          {t('账户充值')}
        </Typography.Title>
        <Button
          icon={<Receipt size={16} />}
          theme='light'
          type='tertiary'
          onClick={onOpenHistory}
        >
          {t('账单')}
        </Button>
      </div>

      {shouldShowSubscription ? (
        <Tabs type='card' activeKey={activeTab} onChange={setActiveTab}>
          <TabPane
            tab={
              <div className='flex items-center gap-2'>
                <Sparkles size={16} />
                {t('订阅套餐')}
              </div>
            }
            itemKey='subscription'
          >
            <div className='py-2'>
              <SubscriptionPlansCard
                t={t}
                loading={subscriptionLoading}
                plans={subscriptionPlans}
                payMethods={payMethods}
                enableOnlineTopUp={enableOnlineTopUp}
                enableStripeTopUp={enableStripeTopUp}
                enableCreemTopUp={enableCreemTopUp}
                billingPreference={billingPreference}
                onChangeBillingPreference={onChangeBillingPreference}
                activeSubscriptions={activeSubscriptions}
                allSubscriptions={allSubscriptions}
                reloadSubscriptionSelf={reloadSubscriptionSelf}
                withCard={false}
              />
            </div>
          </TabPane>
          <TabPane
            tab={
              <div className='flex items-center gap-2'>
                <Wallet size={16} />
                {t('额度充值')}
              </div>
            }
            itemKey='topup'
          >
            <div className='py-2'>{topupContent}</div>
          </TabPane>
        </Tabs>
      ) : (
        topupContent
      )}
    </Card>
  );
};

export default RechargeCard;
