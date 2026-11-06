import React, { useEffect, useState } from 'react';
import { Button, Modal, Typography } from '@douyinfe/semi-ui';
import { IconCopy, IconExternalOpen } from '@douyinfe/semi-icons';
import { useTranslation } from 'react-i18next';
import ResultUrlBlock from './ResultUrlBlock';

const { Text } = Typography;

const ImageResultModal = ({ visible, url, onClose }) => {
  const { t } = useTranslation();
  const [failed, setFailed] = useState(false);

  useEffect(() => {
    setFailed(false);
  }, [url, visible]);

  const copyUrl = () => {
    if (url) navigator.clipboard.writeText(url);
  };

  return (
    <Modal
      title={t('图像预览')}
      visible={visible}
      onCancel={onClose}
      footer={null}
      width={760}
      bodyStyle={{ maxHeight: '76vh', overflow: 'auto', paddingBottom: 24 }}
    >
      {failed ? (
        <div style={{ paddingTop: 4 }}>
          <Text type='tertiary' style={{ display: 'block', marginBottom: 16 }}>
            {t('图像打不开。记录里的地址还在，但这是上游的临时链接，过期后文件就会返回拒绝，预览就没有画面。')}
          </Text>
          <Button
            icon={<IconExternalOpen />}
            onClick={() => window.open(url, '_blank')}
            style={{ marginRight: 8 }}
          >
            {t('在新标签页中打开')}
          </Button>
          <Button icon={<IconCopy />} onClick={copyUrl}>
            {t('复制链接')}
          </Button>
          <ResultUrlBlock url={url} />
        </div>
      ) : (
        <img
          src={url}
          alt=''
          onError={() => setFailed(true)}
          style={{ display: 'block', maxWidth: '100%', margin: '0 auto' }}
        />
      )}
    </Modal>
  );
};

export default ImageResultModal;
