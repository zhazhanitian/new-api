import React, { useEffect, useMemo, useRef, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import {
  Button,
  Input,
  Popconfirm,
  Radio,
  RadioGroup,
  Select,
  Switch,
  Tag,
  Toast,
  Typography,
  TextArea,
} from '@douyinfe/semi-ui';
import { API, showError } from '../../helpers';
import { API_KINDS, runTrialFlow } from '../../helpers/modelTrial';
import NodeEditor, { ensureNodeIds } from '../../components/model-trial/NodeEditor';
import TrialForm from '../../components/model-trial/TrialForm';
import RequestJson from '../../components/model-trial/RequestJson';
import ResultView from '../../components/model-trial/ResultView';

const { Title, Text } = Typography;

const ASIDE_WIDTH_KEY = 'model-trial-aside-width';

function initialAsideWidth() {
  const saved = Number(localStorage.getItem(ASIDE_WIDTH_KEY));
  if (Number.isFinite(saved) && saved >= 360) return saved;
  if (typeof window === 'undefined') return 560;
  return Math.round(Math.min(720, Math.max(520, window.innerWidth * 0.36)));
}

function ColumnSplit({ onResize, onResizeEnd }) {
  const [active, setActive] = useState(false);
  const startDrag = (event) => {
    event.preventDefault();
    setActive(true);
    const previousCursor = document.body.style.cursor;
    const previousSelect = document.body.style.userSelect;
    document.body.style.cursor = 'col-resize';
    document.body.style.userSelect = 'none';
    const move = (ev) => onResize(ev.clientX);
    const up = () => {
      setActive(false);
      document.body.style.cursor = previousCursor;
      document.body.style.userSelect = previousSelect;
      window.removeEventListener('pointermove', move);
      window.removeEventListener('pointerup', up);
      onResizeEnd();
    };
    window.addEventListener('pointermove', move);
    window.addEventListener('pointerup', up);
    onResize(event.clientX);
  };
  return (
    <div
      role='separator'
      aria-orientation='vertical'
      onPointerDown={startDrag}
      style={{
        width: 16,
        flex: 'none',
        cursor: 'col-resize',
        display: 'flex',
        justifyContent: 'center',
        touchAction: 'none',
      }}
    >
      <div
        style={{
          width: active ? 2 : 1,
          height: '100%',
          borderRadius: 1,
          background: active ? 'var(--semi-color-primary)' : 'var(--semi-color-border)',
        }}
      />
    </div>
  );
}

function modelTags(item) {
  return String(item?.tags || '')
    .split(',')
    .map((tag) => tag.trim())
    .filter(Boolean);
}

async function loadAllModels() {
  const all = [];
  let page = 1;
  let total = 0;
  do {
    const res = await API.get(`/api/models/search?keyword=&p=${page}&page_size=100`);
    const data = res.data?.data || {};
    const items = data.items || [];
    total = Number(data.total || 0);
    all.push(...items);
    if (!items.length) break;
    page += 1;
  } while (all.length < total && page <= 40);
  return all;
}

const emptyDraft = {
  id: 0,
  model_name: '',
  api_kind: 'image_generation',
  enabled: true,
  nodes: [],
  examples: [],
};

export default function ModelTrialAdmin() {
  const [searchParams] = useSearchParams();
  const [items, setItems] = useState([]);
  const [draft, setDraft] = useState(emptyDraft);
  const [models, setModels] = useState([]);
  const [modelTag, setModelTag] = useState('');
  const [groups, setGroups] = useState([]);
  const [group, setGroup] = useState('');
  const [values, setValues] = useState({});
  const [trialView, setTrialView] = useState('form');
  const [result, setResult] = useState(null);
  const [streamingText, setStreamingText] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState(false);
  const [testing, setTesting] = useState(false);
  const [exampleTitle, setExampleTitle] = useState('');
  const [exampleDescription, setExampleDescription] = useState('');
  const [asideWidth, setAsideWidth] = useState(initialAsideWidth);
  const asideWidthRef = useRef(asideWidth);
  const splitRef = useRef(null);

  const resizeAside = (clientX) => {
    const rect = splitRef.current?.getBoundingClientRect();
    if (!rect) return;
    const min = 360;
    const max = Math.max(min, rect.width - 480);
    const next = Math.round(Math.min(Math.max(rect.right - clientX, min), max));
    asideWidthRef.current = next;
    setAsideWidth(next);
  };

  const load = async () => {
    const res = await API.get('/api/model-trials/');
    if (!res.data.success) {
      showError(res.data.message);
      return;
    }
    setItems(res.data.data || []);
  };

  useEffect(() => {
    load().catch((err) => showError(err.message));
    loadAllModels()
      .then(setModels)
      .catch(() => {});
    API.get('/api/user/self/groups')
      .then((res) => {
        const names = Object.keys(res.data?.data || {});
        setGroups(names);
        setGroup(names[0] || '');
      })
      .catch(() => {});
  }, []);

  useEffect(() => {
    const modelName = searchParams.get('model');
    if (!modelName || !items.length) return;
    const found = items.find((item) => item.model_name === modelName);
    if (found) setDraft({ ...found, nodes: ensureNodeIds(found.nodes) });
    else setDraft({ ...emptyDraft, model_name: modelName });
  }, [searchParams, items]);

  const modelTagOptions = useMemo(() => {
    const counts = new Map();
    models.forEach((item) => {
      modelTags(item).forEach((tag) => counts.set(tag, (counts.get(tag) || 0) + 1));
    });
    return [...counts.entries()].sort(
      (a, b) => b[1] - a[1] || a[0].localeCompare(b[0], 'zh'),
    );
  }, [models]);

  const modelOptions = useMemo(
    () =>
      models
        .filter((item) => !modelTag || modelTags(item).includes(modelTag))
        .map((item) => ({
          value: item.model_name,
          label: item.model_name,
          tags: modelTags(item),
        })),
    [models, modelTag],
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

  const save = async () => {
    const payload = {
      model_name: draft.model_name,
      api_kind: draft.api_kind,
      enabled: draft.enabled,
      nodes: draft.nodes,
    };
    const res = draft.id
      ? await API.put(`/api/model-trials/${draft.id}`, payload)
      : await API.post('/api/model-trials/', payload);
    if (!res.data.success) {
      showError(res.data.message);
      return;
    }
    setDraft(res.data.data);
    Toast.success('已保存');
    load();
  };

  const test = async (nextValues) => {
    const payload = nextValues || values;
    setTesting(true);
    setError('');
    setSuccess(false);
    setResult(null);
    setStreamingText('');
    try {
      await runTrialFlow({
        model: draft.model_name,
        apiKind: draft.api_kind,
        group,
        values: payload,
        nodes: draft.nodes,
        onUpdate: (next) => {
          setResult(next.result);
          setStreamingText(next.streamingText || '');
          setError(next.error || '');
          setSuccess(!!next.success);
        },
      });
    } catch (err) {
      setError(err.message);
      setSuccess(false);
    } finally {
      setTesting(false);
    }
  };

  const saveExample = async () => {
    if (!draft.id) {
      showError('请先保存配置');
      return;
    }
    const res = await API.post(`/api/model-trials/${draft.id}/examples`, {
      title: exampleTitle,
      description: exampleDescription,
      values,
    });
    if (!res.data.success) {
      showError(res.data.message);
      return;
    }
    setDraft(res.data.data);
    setExampleTitle('');
    setExampleDescription('');
    Toast.success('示例已保存');
    load();
  };

  const applyExample = (example) => {
    if (example.stale) {
      Toast.warning('字段已经改过，这个示例对不上，请重新测试后再保存');
      return;
    }
    setValues(example.values || {});
    setResult(null);
    setStreamingText('');
    setError('');
    setSuccess(false);
  };

  const removeExample = async (exampleId) => {
    const res = await API.delete(`/api/model-trials/${draft.id}/examples/${exampleId}`);
    if (!res.data.success) {
      showError(res.data.message);
      return;
    }
    setDraft(res.data.data);
    load();
  };

  const openDraft = (item) => {
    setDraft({ ...item, nodes: ensureNodeIds(item.nodes) });
    setValues({});
    setResult(null);
    setStreamingText('');
    setError('');
    setSuccess(false);
  };

  const clearDraft = () => {
    setDraft({
      ...emptyDraft,
      model_name: draft.model_name,
      api_kind: draft.api_kind,
      enabled: draft.enabled,
    });
    setValues({});
    setResult(null);
    setStreamingText('');
    setError('');
    setSuccess(false);
    setExampleTitle('');
    setExampleDescription('');
  };

  return (
    <div
      className='px-6'
      style={{
        marginTop: 72,
        height: 'calc(100vh - 72px)',
        display: 'flex',
        flexDirection: 'column',
        overflow: 'hidden',
        paddingBottom: 16,
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', gap: 16, marginBottom: 12, flex: 'none' }}>
        <Title heading={4} style={{ margin: 0 }}>模型体验配置</Title>
        <Button theme='solid' onClick={save}>保存配置</Button>
      </div>

      <div
        style={{
          display: 'flex',
          flexWrap: 'wrap',
          gap: 12,
          alignItems: 'center',
          marginBottom: 12,
          flex: 'none',
          padding: '14px 16px',
          border: '1px solid var(--semi-color-border)',
          borderRadius: 12,
          background: 'var(--semi-color-bg-1)',
        }}
      >
        <Select
          filter
          placeholder='已有配置'
          style={{ width: 220 }}
          value={draft.id || undefined}
          optionList={items.map((item) => ({ value: item.id, label: item.model_name }))}
          onChange={(id) => {
            const found = items.find((item) => item.id === id);
            if (found) openDraft(found);
          }}
        />
        <Select
          filter
          searchPosition='dropdown'
          searchPlaceholder='搜索模型名称'
          placeholder='选择模型'
          style={{ width: 280 }}
          dropdownMatchSelectWidth={false}
          dropdownStyle={{ width: 420 }}
          value={draft.model_name || undefined}
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
          onChange={(model_name) => setDraft({ ...draft, model_name })}
        />
        <Select
          style={{ width: 180 }}
          value={draft.api_kind}
          optionList={API_KINDS}
          onChange={(api_kind) => setDraft({ ...draft, api_kind })}
        />
        <span style={{ display: 'inline-flex', alignItems: 'center', gap: 8, height: 32 }}>
          <span>启用</span>
          <Switch checked={!!draft.enabled} onChange={(enabled) => setDraft({ ...draft, enabled })} />
        </span>
        <Button onClick={() => openDraft({ ...emptyDraft })}>新建配置</Button>
        <Popconfirm
          position='bottomRight'
          title='清空当前编辑区？'
          content='字段会全部去掉，已保存的配置不会删除。'
          onConfirm={clearDraft}
        >
          <Button type='tertiary' disabled={!draft.id && !(draft.nodes || []).length}>
            清空
          </Button>
        </Popconfirm>
      </div>

      <div
        ref={splitRef}
        style={{ flex: 1, minHeight: 0, display: 'flex', alignItems: 'stretch' }}
      >
        <section style={{ flex: 1, minWidth: 0, minHeight: 0, overflow: 'auto', paddingRight: 8 }}>
          <NodeEditor
            apiKind={draft.api_kind}
            nodes={draft.nodes}
            onChange={(nodes) => setDraft({ ...draft, nodes })}
          />
        </section>

        <ColumnSplit
          onResize={resizeAside}
          onResizeEnd={() => localStorage.setItem(ASIDE_WIDTH_KEY, String(asideWidthRef.current))}
        />

        <aside
          style={{
            width: asideWidth,
            flex: 'none',
            minHeight: 0,
            display: 'flex',
            flexDirection: 'column',
            border: '1px solid var(--semi-color-border)',
            borderRadius: 12,
            background: 'var(--semi-color-bg-1)',
            overflow: 'hidden',
          }}
        >
          <div style={{ flex: 1, minHeight: 0, overflow: 'auto', padding: 16 }}>
            <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', gap: 8 }}>
              <Text strong>试调</Text>
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
            <Text type='tertiary' size='small' style={{ display: 'block', margin: '4px 0 12px' }}>
              {trialView === 'json' ? '和正式请求使用同一套组装，可复制。' : '表单顺序与左侧卡片一致。'}
            </Text>
            <Select
              style={{ width: '100%', marginBottom: 16 }}
              placeholder='分组'
              value={group || undefined}
              optionList={groups.map((item) => ({ value: item, label: item }))}
              onChange={setGroup}
            />
            {(draft.examples || []).length > 0 && (
              <div style={{ marginBottom: 16 }}>
                <Text strong>已保存示例</Text>
                <Text type='tertiary' size='small' style={{ display: 'block', margin: '4px 0 8px' }}>
                  点一下只填入当时的参数，确认后再点测试。
                </Text>
                <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
                  {(draft.examples || []).map((example) => (
                    <div key={example.id}>
                      <div style={{ display: 'flex', alignItems: 'flex-start', gap: 4 }}>
                        <Button
                          theme='light'
                          size='small'
                          disabled={example.stale || testing}
                          onClick={() => applyExample(example)}
                          style={{ flex: 1, justifyContent: 'flex-start' }}
                        >
                          {example.title}
                          {example.stale ? '（已失效）' : ''}
                        </Button>
                        <Button
                          size='small'
                          type='danger'
                          theme='borderless'
                          onClick={() => removeExample(example.id)}
                        >
                          删除
                        </Button>
                      </div>
                      {example.description && (
                        <Text type='tertiary' size='small' style={{ display: 'block', marginTop: 2 }}>
                          {example.description}
                        </Text>
                      )}
                    </div>
                  ))}
                </div>
              </div>
            )}
            {trialView === 'json' ? (
              <RequestJson
                modelName={draft.model_name}
                apiKind={draft.api_kind}
                nodes={draft.nodes}
                values={values}
              />
            ) : (
              <TrialForm nodes={draft.nodes} values={values} onChange={setValues} />
            )}
            <Button theme='solid' loading={testing} onClick={() => test()} disabled={!draft.model_name}>
              测试
            </Button>
            <div style={{ marginTop: 16, paddingTop: 16, borderTop: '1px solid var(--semi-color-border)' }}>
              <Text strong>结果</Text>
              <div style={{ marginTop: 12 }}>
                <ResultView result={result} error={error} streamingText={streamingText} />
              </div>
              <Input
                style={{ marginTop: 12 }}
                placeholder='示例标题'
                value={exampleTitle}
                onChange={setExampleTitle}
              />
              <TextArea
                style={{ marginTop: 8 }}
                placeholder='示例描述'
                value={exampleDescription}
                autosize={{ minRows: 1, maxRows: 3 }}
                onChange={setExampleDescription}
              />
              <Button style={{ marginTop: 8 }} disabled={!success || !exampleTitle} onClick={saveExample}>
                保存为示例
              </Button>
            </div>
          </div>
        </aside>
      </div>
    </div>
  );
}
