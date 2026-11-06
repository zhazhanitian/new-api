import React, { useEffect, useState, useRef } from 'react';
import { Banner, Button, Form, Row, Col, Spin, Typography } from '@douyinfe/semi-ui';
import { API, removeTrailingSlash, showError, showSuccess } from '../../../helpers';
import { useTranslation } from 'react-i18next';

const { Text } = Typography;

export default function SettingsPaymentGatewayPayCenter(props) {
  const { t } = useTranslation();
  const sectionTitle = props.hideSectionTitle ? undefined : t('支付中心设置');
  const [loading, setLoading] = useState(false);
  const [inputs, setInputs] = useState({
    PayCenterEnabled: false,
    PayCenterBaseURL: '',
    PayCenterApiKey: '',
    PayCenterNotifyKey: '',
    PayCenterProductCode: '',
    PayCenterUnitPrice: 7.3,
    PayCenterMinTopUp: 1,
  });
  const formApiRef = useRef(null);

  useEffect(() => {
    if (props.options && formApiRef.current) {
      const currentInputs = {
        PayCenterEnabled: props.options.PayCenterEnabled === true,
        PayCenterBaseURL: props.options.PayCenterBaseURL || '',
        PayCenterApiKey: '',
        PayCenterNotifyKey: '',
        PayCenterProductCode: props.options.PayCenterProductCode || '',
        PayCenterUnitPrice:
          props.options.PayCenterUnitPrice !== undefined
            ? parseFloat(props.options.PayCenterUnitPrice)
            : 7.3,
        PayCenterMinTopUp:
          props.options.PayCenterMinTopUp !== undefined
            ? parseFloat(props.options.PayCenterMinTopUp)
            : 1,
      };
      setInputs(currentInputs);
      formApiRef.current.setValues(currentInputs);
    }
  }, [props.options]);

  const handleFormChange = (values) => {
    setInputs(values);
  };

  const callbackBase = removeTrailingSlash(
    props.options?.CustomCallbackAddress || props.options?.ServerAddress || '',
  );
  const notifyUrl = callbackBase ? `${callbackBase}/api/paycenter/notify` : '';

  const submit = async () => {
    setLoading(true);
    try {
      const options = [
        {
          key: 'PayCenterEnabled',
          value: inputs.PayCenterEnabled === true,
        },
        {
          key: 'PayCenterBaseURL',
          value: removeTrailingSlash(inputs.PayCenterBaseURL || ''),
        },
        {
          key: 'PayCenterProductCode',
          value: (inputs.PayCenterProductCode || '').trim(),
        },
        {
          key: 'PayCenterUnitPrice',
          value: String(inputs.PayCenterUnitPrice ?? ''),
        },
        {
          key: 'PayCenterMinTopUp',
          value: String(inputs.PayCenterMinTopUp ?? ''),
        },
      ];
      if (inputs.PayCenterApiKey) {
        options.push({
          key: 'PayCenterApiKey',
          value: inputs.PayCenterApiKey.trim(),
        });
      }
      if (inputs.PayCenterNotifyKey) {
        options.push({
          key: 'PayCenterNotifyKey',
          value: inputs.PayCenterNotifyKey.trim(),
        });
      }

      const results = await Promise.all(
        options.map((opt) =>
          API.put('/api/option/', {
            key: opt.key,
            value: opt.value,
          }),
        ),
      );
      const errorResults = results.filter((res) => !res.data.success);
      if (errorResults.length > 0) {
        errorResults.forEach((res) => {
          showError(res.data.message);
        });
      } else {
        showSuccess(t('更新成功'));
        props.refresh && props.refresh();
      }
    } catch (error) {
      showError(t('更新失败'));
    }
    setLoading(false);
  };

  return (
    <Spin spinning={loading}>
      <Form
        initValues={inputs}
        onValueChange={handleFormChange}
        getFormApi={(api) => (formApiRef.current = api)}
      >
        <Form.Section text={sectionTitle}>
          <Banner
            type='info'
            description={t(
              '回调地址要填到支付中心的产品配置里，下单请求不会携带通知地址。',
            )}
            style={{ marginBottom: 16 }}
          />
          {notifyUrl ? (
            <Banner
              type='success'
              description={
                <Text>
                  {t('回调地址')}：{notifyUrl}
                </Text>
              }
              style={{ marginBottom: 16 }}
            />
          ) : (
            <Banner
              type='warning'
              description={t(
                '服务器地址为空。请先在通用设置填写服务器地址，或填写自定义回调地址。',
              )}
              style={{ marginBottom: 16 }}
            />
          )}
          <Row gutter={{ xs: 8, sm: 16, md: 24, lg: 24, xl: 24, xxl: 24 }}>
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.Switch
                field='PayCenterEnabled'
                label={t('启用支付中心')}
              />
            </Col>
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.Input
                field='PayCenterBaseURL'
                label={t('支付中心地址')}
                placeholder={t('例如：https://pay.example.com')}
              />
            </Col>
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.Input
                field='PayCenterProductCode'
                label={t('产品编码')}
                placeholder='MY_MAPI'
              />
            </Col>
          </Row>
          <Row
            gutter={{ xs: 8, sm: 16, md: 24, lg: 24, xl: 24, xxl: 24 }}
            style={{ marginTop: 16 }}
          >
            <Col xs={24} sm={24} md={12} lg={12} xl={12}>
              <Form.Input
                field='PayCenterApiKey'
                label={t('API Key')}
                placeholder={t('敏感信息不会发送到前端显示，留空则不修改')}
                type='password'
              />
            </Col>
            <Col xs={24} sm={24} md={12} lg={12} xl={12}>
              <Form.Input
                field='PayCenterNotifyKey'
                label={t('通知秘钥')}
                placeholder={t('敏感信息不会发送到前端显示，留空则不修改')}
                type='password'
              />
            </Col>
          </Row>
          <Row
            gutter={{ xs: 8, sm: 16, md: 24, lg: 24, xl: 24, xxl: 24 }}
            style={{ marginTop: 16 }}
          >
            <Col xs={24} sm={24} md={12} lg={12} xl={12}>
              <Form.InputNumber
                field='PayCenterUnitPrice'
                precision={2}
                label={t('充值价格（元/美元）')}
                placeholder={t('例如：7，就是 7 元/美元')}
              />
            </Col>
            <Col xs={24} sm={24} md={12} lg={12} xl={12}>
              <Form.InputNumber
                field='PayCenterMinTopUp'
                precision={0}
                label={t('最低充值美元数量')}
                placeholder={t('例如：1，就是最低充值 1 美元')}
              />
            </Col>
          </Row>
          <Button onClick={submit} style={{ marginTop: 16 }}>
            {t('更新支付中心设置')}
          </Button>
        </Form.Section>
      </Form>
    </Spin>
  );
}
