import React, { useEffect, useMemo, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import { Button, Radio, RadioGroup, Select, Tag, Toast, Typography } from '@douyinfe/semi-ui';
import { API, showError } from '../../helpers';
import { API_KINDS, runTrialFlow } from '../../helpers/modelTrial';
import TrialForm from '../../components/model-trial/TrialForm';
import RequestJson from '../../components/model-trial/RequestJson';
import ResultView from '../../components/model-trial/ResultView';

const { Title, Text, Paragraph } = Typography;

function trialTags(item) {
  return String(item?.tags || '')
    .split(',')
    .map((tag) => tag.trim())
    .filter(Boolean);
}

export default function ModelTrialPage() {
  const [searchParams] = useSearchParams();
  const requestedModel = searchParams.get('model') || '';
  const [trials, setTrials] = useState([]);
  const [modelTag, setModelTag] = useState('');
  const [currentId, setCurrentId] = useState(0);
  const [groups, setGroups] = useState([]);
  const [group, setGroup] = useState('');
  const [values, setValues] = useState({});
  const [result, setResult] = useState(null);
  const [streamingText, setStreamingText] = useState('');
  const [error, setError] = useState('');
  const [running, setRunning] = useState(false);
  const [trialView, setTrialView] = useState('form');

  useEffect(() => {
    let cancelled = false;
    API.get('/api/user/model-trials')
      .then((res) => {
        if (cancelled) return;
        const list = res.data?.data || [];
        setTrials(list);
        const found = requestedModel
          ? list.find((item) => item.model_name === requestedModel)
          : list[0];
        setCurrentId(found?.id || 0);
        setValues({});
        setResult(null);
        setError('');
        setStreamingText('');
      })
      .catch((err) => {
        if (!cancelled) showError(err.message);
      });
    return () => {
      cancelled = true;
    };
  }, [requestedModel]);

  useEffect(() => {
    API.get('/api/user/self/groups')
      .then((res) => {
        const names = Object.keys(res.data?.data || {});
        setGroups(names);
        setGroup(names[0] || '');
      })
      .catch(() => {});
  }, []);

  const current = trials.find((item) => item.id === currentId);
  const kindLabel = API_KINDS.find((item) => item.value === current?.api_kind)?.label;

  const modelTagOptions = useMemo(() => {
    const counts = new Map();
    trials.forEach((item) => {
      trialTags(item).forEach((tag) => counts.set(tag, (counts.get(tag) || 0) + 1));
    });
    return [...counts.entries()].sort(
      (a, b) => b[1] - a[1] || a[0].localeCompare(b[0], 'zh'),
    );
  }, [trials]);

  const modelOptions = useMemo(
    () =>
      trials
        .filter((item) => !modelTag || trialTags(item).includes(modelTag))
        .map((item) => ({
          value: item.id,
          label: item.model_name,
          tags: trialTags(item),
        })),
    [trials, modelTag],
  );

  const renderModelOption = (props) => {
    const { label, tags = [], focused, selected, style, className, onClick, onMouseEnter } = props;
    return (
      <div
        style={{
          ...style,
          display: 'flex',
          alignItems: 'center',
          gap: 8,
          padding: '6px 12px',
          background: focused
            ? 'var(--semi-color-fill-0)'
            : selected
              ? 'var(--semi-color-primary-light-default)'
              : undefined,
        }}
        className={className}
        onClick={onClick}
        onMouseEnter={onMouseEnter}
      >
        <span
          style={{
            flex: 1,
            minWidth: 0,
            overflow: 'hidden',
            textOverflow: 'ellipsis',
            whiteSpace: 'nowrap',
          }}
        >
          {label}
        </span>
        <span style={{ display: 'inline-flex', gap: 4, flex: 'none' }}>
          {tags.slice(0, 3).map((tag) => (
            <Tag key={tag} size='small' color='blue'>
              {tag}
            </Tag>
          ))}
        </span>
      </div>
    );
  };

  const run = async (nextValues) => {
    if (!current) return;
    const payload = nextValues || values;
    setRunning(true);
    setError('');
    setResult(null);
    setStreamingText('');
    try {
      await runTrialFlow({
        model: current.model_name,
        apiKind: current.api_kind,
        group,
        values: payload,
        nodes: current.nodes,
        onUpdate: (next) => {
          setResult(next.result);
          setStreamingText(next.streamingText || '');
          setError(next.error || '');
        },
      });
    } catch (err) {
      setError(err.message);
    } finally {
      setRunning(false);
    }
  };

  const applyExample = (example) => {
    if (example.stale) {
      Toast.warning('这个示例已失效，请联系管理员重新测试');
      return;
    }
    setValues(example.values || {});
    setResult(null);
    setStreamingText('');
    setError('');
  };

  return (
    <div className='mt-[60px] px-4 pb-8'>
      <Title heading={4}>模型体验</Title>
      {!trials.length && <Text type='tertiary'>还没有可体验的模型</Text>}
      {!!requestedModel &&
        !!trials.length &&
        !trials.some((item) => item.model_name === requestedModel) && (
          <Text type='tertiary' style={{ display: 'block', marginTop: 12 }}>
            {requestedModel} 还没有配置在线体验
          </Text>
        )}
      {!!trials.length && (
        <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 24, marginTop: 16 }}>
          <div>
            <Select
              filter
              searchPosition='dropdown'
              searchPlaceholder='搜索模型'
              style={{ width: '100%' }}
              placeholder='搜索模型'
              value={currentId || undefined}
              optionList={modelOptions}
              renderOptionItem={renderModelOption}
              outerTopSlot={
                modelTagOptions.length ? (
                  <div
                    style={{
                      display: 'flex',
                      flexWrap: 'wrap',
                      gap: 6,
                      padding: '8px 12px 4px',
                      maxHeight: 108,
                      overflow: 'auto',
                    }}
                    onMouseDown={(event) => event.preventDefault()}
                  >
                    <Button
                      size='small'
                      theme={modelTag ? 'light' : 'solid'}
                      type={modelTag ? 'tertiary' : 'primary'}
                      onClick={() => setModelTag('')}
                    >
                      全部
                    </Button>
                    {modelTagOptions.map(([tag, count]) => (
                      <Button
                        key={tag}
                        size='small'
                        theme={modelTag === tag ? 'solid' : 'light'}
                        type={modelTag === tag ? 'primary' : 'tertiary'}
                        onClick={() => setModelTag(modelTag === tag ? '' : tag)}
                      >
                        {tag} {count}
                      </Button>
                    ))}
                  </div>
                ) : null
              }
              onChange={(id) => {
                setCurrentId(id);
                setValues({});
                setResult(null);
                setError('');
                setStreamingText('');
              }}
            />
            {current && (
              <div style={{ margin: '12px 0' }}>
                <Tag>{kindLabel || current.api_kind}</Tag>
              </div>
            )}
            <Select
              style={{ width: '100%', marginBottom: 12 }}
              placeholder='分组'
              value={group || undefined}
              optionList={groups.map((item) => ({ value: item, label: item }))}
              onChange={setGroup}
            />
            {(current?.examples || []).length > 0 && (
              <div style={{ marginBottom: 16 }}>
                <Text strong>调用示例</Text>
                <Text type='tertiary' size='small' style={{ display: 'block', marginTop: 4 }}>
                  点一下只填入参数，确认后再点运行。
                </Text>
                {(current.examples || []).map((example) => (
                  <div key={example.id} style={{ marginTop: 8 }}>
                    <Button
                      theme='light'
                      disabled={example.stale || running}
                      onClick={() => applyExample(example)}
                    >
                      {example.title}
                      {example.stale ? '（已失效）' : ''}
                    </Button>
                    {example.description && (
                      <Paragraph type='tertiary' style={{ margin: '4px 0 0' }}>
                        {example.description}
                      </Paragraph>
                    )}
                  </div>
                ))}
              </div>
            )}
            <div style={{ display: 'flex', justifyContent: 'flex-end', marginBottom: 12 }}>
              <RadioGroup
                type='button'
                size='small'
                value={trialView}
                onChange={(e) => setTrialView(e.target.value)}
              >
                <Radio value='form'>表单</Radio>
                <Radio value='json'>JSON</Radio>
              </RadioGroup>
            </div>
            {trialView === 'json' ? (
              <RequestJson
                modelName={current?.model_name}
                apiKind={current?.api_kind}
                nodes={current?.nodes || []}
                values={values}
              />
            ) : (
              <TrialForm nodes={current?.nodes || []} values={values} onChange={setValues} />
            )}
            <Button theme='solid' loading={running} disabled={!current} onClick={() => run()}>
              运行
            </Button>
          </div>
          <div>
            <Text strong>结果</Text>
            <div style={{ marginTop: 12 }}>
              <ResultView result={result} error={error} streamingText={streamingText} />
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
