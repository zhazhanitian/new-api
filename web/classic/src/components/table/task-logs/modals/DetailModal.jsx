import React from 'react';
import { Modal, Typography } from '@douyinfe/semi-ui';
import { useTranslation } from 'react-i18next';
import JsonCodeBlock from '../../../common/JsonCodeBlock';

const { Text } = Typography;

function payloadProps(payload) {
  if (payload == null || payload === '') {
    return null;
  }
  if (typeof payload === 'string') {
    return { text: payload };
  }
  return { value: payload };
}

function userInput(record) {
  if (record?.user_request) {
    return record.user_request;
  }
  const input = record?.properties?.input;
  if (typeof input === 'string' && input.trim()) {
    return input;
  }
  return null;
}

// 和原先点击任务 ID 弹出的是同一份记录。上游原文留在「查看原文」里。
function taskDetailText(record) {
  if (!record) {
    return '';
  }
  const {
    user_request: _userRequest,
    upstream_request: _upstreamRequest,
    upstream_response: _upstreamResponse,
    key: _key,
    timestamp2string: _timestamp,
    ...detail
  } = record;
  return JSON.stringify(detail, null, 2);
}

const DetailModal = ({ detailRecord, setDetailRecord }) => {
  const { t } = useTranslation();
  const visible = !!detailRecord;
  const inputProps = payloadProps(userInput(detailRecord));

  return (
    <Modal
      title={t('任务详情')}
      visible={visible}
      onCancel={() => setDetailRecord(null)}
      footer={null}
      width={800}
      bodyStyle={{ maxHeight: '72vh', overflow: 'auto', paddingBottom: 24 }}
    >
      <div className='flex flex-col gap-5'>
        <div>
          <Text strong style={{ display: 'block', marginBottom: 8 }}>
            {t('用户传入')}
          </Text>
          {inputProps ? (
            <JsonCodeBlock {...inputProps} />
          ) : (
            <Text type='tertiary'>{t('未记录')}</Text>
          )}
        </div>
        <div>
          <Text strong style={{ display: 'block', marginBottom: 8 }}>
            {t('任务详情')}
          </Text>
          <JsonCodeBlock text={taskDetailText(detailRecord)} />
        </div>
      </div>
    </Modal>
  );
};

export default DetailModal;
