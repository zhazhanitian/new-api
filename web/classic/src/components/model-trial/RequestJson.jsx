import React, { useEffect, useState } from 'react';
import { Typography } from '@douyinfe/semi-ui';
import { API } from '../../helpers';
import JsonCodeBlock from '../common/JsonCodeBlock';

const { Text } = Typography;

export default function RequestJson({ modelName, apiKind, nodes, values }) {
  const [payload, setPayload] = useState(null);
  const [error, setError] = useState('');
  const signature = JSON.stringify({ modelName, apiKind, nodes, values });

  useEffect(() => {
    let cancelled = false;
    const timer = setTimeout(() => {
      API.post('/api/user/model-trial/preview', {
        model_name: modelName,
        api_kind: apiKind,
        nodes: nodes || [],
        values: values || {},
      })
        .then((res) => {
          if (cancelled) return;
          if (!res.data?.success) {
            setPayload(null);
            setError(res.data?.message || '暂时无法生成请求');
            return;
          }
          setError('');
          setPayload(res.data.data);
        })
        .catch((err) => {
          if (cancelled) return;
          setPayload(null);
          setError(err?.response?.data?.message || err.message || '暂时无法生成请求');
        });
    }, 250);
    return () => {
      cancelled = true;
      clearTimeout(timer);
    };
  }, [signature, modelName, apiKind, nodes, values]);

  if (error) {
    return <Text type='danger'>{error}</Text>;
  }
  if (!payload) {
    return <Text type='tertiary'>正在生成请求...</Text>;
  }
  return (
    <div style={{ marginBottom: 24 }}>
      <Text type='tertiary' size='small' style={{ display: 'block', marginBottom: 8 }}>
        {payload.method} {payload.path}
      </Text>
      <JsonCodeBlock value={payload.body} />
    </div>
  );
}
