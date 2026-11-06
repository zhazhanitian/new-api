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

function RequestBlock({ title, payload, emptyText }) {
  const props = payloadProps(payload);
  return (
    <div>
      <Text strong style={{ display: 'block', marginBottom: 8 }}>
        {title}
      </Text>
      {props ? (
        <JsonCodeBlock {...props} />
      ) : (
        <Text type='tertiary'>{emptyText}</Text>
      )}
    </div>
  );
}

const RequestModal = ({ requestRecord, setRequestRecord }) => {
  const { t } = useTranslation();
  const visible = !!requestRecord;

  return (
    <Modal
      title={t('原始记录')}
      visible={visible}
      onCancel={() => setRequestRecord(null)}
      footer={null}
      width={880}
      bodyStyle={{ maxHeight: '72vh', overflow: 'auto', paddingBottom: 24 }}
    >
      <div className='flex flex-col gap-5'>
        {!requestRecord?.user_request &&
        !requestRecord?.upstream_request &&
        !requestRecord?.upstream_response ? (
          <Text type='tertiary'>
            {t('这条是功能上线前的任务，没有留下原始记录。新提交的任务可以在这里看到用户输入、上游请求和上游响应。')}
          </Text>
        ) : (
          <>
            <RequestBlock
              title={t('用户原始输入')}
              payload={requestRecord?.user_request}
              emptyText={t('未记录')}
            />
            <RequestBlock
              title={t('上游请求')}
              payload={requestRecord?.upstream_request}
              emptyText={t('未记录')}
            />
            <RequestBlock
              title={t('上游响应')}
              payload={requestRecord?.upstream_response}
              emptyText={t('未记录')}
            />
            <Text type='tertiary' size='small'>
              {t('图片和文件内容只保留长度，提示词、参数和图片地址会原样保留。')}
            </Text>
          </>
        )}
      </div>
    </Modal>
  );
};

export default RequestModal;
