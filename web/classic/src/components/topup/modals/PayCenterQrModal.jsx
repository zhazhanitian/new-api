import React, { useEffect, useRef } from 'react';
import { Modal, Typography } from '@douyinfe/semi-ui';
import { QRCodeSVG } from 'qrcode.react';
import { API, showError, showInfo } from '../../../helpers';

const { Text } = Typography;
const POLL_INTERVAL_MS = 2000;
const EXPIRE_MS = 30 * 60 * 1000;

const PayCenterQrModal = ({ t, order, onCancel, onSuccess }) => {
  const onCancelRef = useRef(onCancel);
  const onSuccessRef = useRef(onSuccess);
  onCancelRef.current = onCancel;
  onSuccessRef.current = onSuccess;

  useEffect(() => {
    if (!order?.tradeNo) {
      return undefined;
    }
    let stopped = false;
    const started = Date.now();
    const tradeNo = order.tradeNo;

    const finish = (kind) => {
      if (stopped) return;
      stopped = true;
      if (kind === 'success') {
        onSuccessRef.current();
        return;
      }
      if (kind === 'expired') {
        showInfo(t('二维码已过期，请重新下单'));
      } else {
        showError(t('支付未完成'));
      }
      onCancelRef.current();
    };

    const tick = async () => {
      if (stopped) return;
      if (Date.now() - started >= EXPIRE_MS) {
        finish('expired');
        return;
      }
      try {
        const res = await API.get(
          `/api/user/paycenter/orders/${encodeURIComponent(tradeNo)}`,
        );
        if (stopped) return;
        const payload = res?.data;
        if (payload?.message !== 'success') {
          return;
        }
        const status = payload.data?.status;
        if (status === 'success') {
          finish('success');
          return;
        }
        if (
          status === 'failed' ||
          status === 'expired' ||
          payload.data?.stop
        ) {
          finish('failed');
        }
      } catch (error) {
        // 查单失败时继续轮询，等待回调或下一次补偿。
      }
    };

    tick();
    const timer = setInterval(tick, POLL_INTERVAL_MS);
    return () => {
      stopped = true;
      clearInterval(timer);
    };
  }, [order?.tradeNo, t]);

  return (
    <Modal
      title={t('微信扫码支付')}
      visible={!!order?.codeUrl}
      onCancel={onCancel}
      footer={null}
      maskClosable={false}
      centered
      width={360}
    >
      <div className='flex flex-col items-center gap-3 py-2'>
        {order?.codeUrl ? (
          <QRCodeSVG value={order.codeUrl} size={220} />
        ) : null}
        <Text type='secondary'>{t('请使用微信扫描二维码完成支付')}</Text>
      </div>
    </Modal>
  );
};

export default PayCenterQrModal;
