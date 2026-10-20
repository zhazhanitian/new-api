import { API } from './api';
import { getUserIdFromLocalStorage } from './utils';

export const API_KINDS = [
  { value: 'chat_completions', label: '对话补全' },
  { value: 'responses', label: 'Responses' },
  { value: 'image_generation', label: '图片同步' },
  { value: 'image_task', label: '图片异步' },
  { value: 'video', label: '视频' },
  { value: 'music', label: '音乐' },
  { value: 'threed', label: '3D' },
];

export const ASYNC_KINDS = ['image_task', 'video', 'music', 'threed'];

export const WIDGETS = [
  { value: 'text', label: '单行文本' },
  { value: 'textarea', label: '多行文本' },
  { value: 'number', label: '数字' },
  { value: 'integer', label: '整数' },
  { value: 'slider', label: '拖动条' },
  { value: 'boolean', label: '开关' },
  { value: 'select', label: '下拉' },
  { value: 'image_url', label: '图片地址' },
  { value: 'image_urls', label: '图片地址列表' },
  { value: 'video_url', label: '视频地址' },
  { value: 'audio_url', label: '音频地址' },
  { value: 'file_url', label: '文件地址' },
  { value: 'messages', label: '对话消息' },
  { value: 'input', label: 'Responses 输入' },
  { value: 'json', label: 'JSON' },
];

export function newNodeId() {
  if (typeof crypto !== 'undefined' && crypto.randomUUID) {
    return crypto.randomUUID();
  }
  return `n_${Math.random().toString(16).slice(2)}`;
}

export function blankNode(type) {
  const id = newNodeId();
  if (type === 'select') {
    return {
      id,
      type,
      label: '下拉',
      help: '',
      submit: true,
      required: false,
      key: '',
      allow_custom: false,
      options: [{ value: '', label: '选项', help: '', nodes: [] }],
    };
  }
  if (type === 'branch') {
    return {
      id,
      type,
      label: '分支',
      help: '',
      submit: false,
      required: true,
      options: [
        { value: 'a', label: '选项一', help: '', nodes: [] },
        { value: 'b', label: '选项二', help: '', nodes: [] },
      ],
    };
  }
  if (type === 'object') {
    return {
      id,
      type,
      label: '参数对象',
      help: '',
      submit: true,
      key: '',
      nodes: [],
    };
  }
  return {
    id,
    type: 'field',
    label: '字段',
    help: '',
    submit: true,
    required: false,
    key: '',
    widget: 'text',
  };
}

export function readStreamFlag(nodes, values) {
  let streaming = false;
  const walk = (list) => {
    (list || []).forEach((node) => {
      if (
        node.type === 'field' &&
        (node.key === 'stream' || node.key === '/stream') &&
        values?.[node.id] === true
      ) {
        streaming = true;
      }
      if (node.type === 'object') walk(node.nodes);
      if (node.type === 'select' || node.type === 'branch') {
        const selected = values?.[node.id];
        const option = (node.options || []).find((item) => item.value === selected);
        if (option) walk(option.nodes);
      }
    });
  };
  walk(nodes);
  return streaming;
}

export function extractTaskId(payload) {
  if (!payload || typeof payload !== 'object') return '';
  if (typeof payload.task_id === 'string' && payload.task_id) return payload.task_id;
  const data = payload.data;
  if (data && !Array.isArray(data)) {
    if (typeof data.task_id === 'string' && data.task_id) return data.task_id;
    if (typeof data.id === 'string' && data.id) return data.id;
  }
  if (Array.isArray(data) && data[0]?.task_id) return data[0].task_id;
  return '';
}

function trialHeaders() {
  return {
    'Content-Type': 'application/json',
    'New-Api-User': getUserIdFromLocalStorage(),
  };
}

export async function submitModelTrial({ model, group, values, stream, onDelta }) {
  const url = `/pg/trial/${encodeURIComponent(model)}`;
  const body = JSON.stringify({ group, values });
  if (stream) {
    const response = await fetch(url, { method: 'POST', headers: trialHeaders(), body });
    const proofId = response.headers.get('X-Model-Trial-Proof') || '';
    if (!response.ok) {
      const text = await response.text();
      throw new Error(text || '调用失败');
    }
    const reader = response.body?.getReader();
    const decoder = new TextDecoder();
    let buffer = '';
    let text = '';
    if (!reader) return { proofId, payload: { text }, stream: true };
    while (true) {
      const { done, value } = await reader.read();
      if (done) break;
      buffer += decoder.decode(value, { stream: true });
      const lines = buffer.split('\n');
      buffer = lines.pop() || '';
      lines.forEach((line) => {
        const trimmed = line.trim();
        if (!trimmed.startsWith('data:')) return;
        const data = trimmed.slice(5).trim();
        if (!data || data === '[DONE]') return;
        try {
          const json = JSON.parse(data);
          const delta =
            json.choices?.[0]?.delta?.content ||
            json.choices?.[0]?.delta?.reasoning_content ||
            json.delta ||
            '';
          if (typeof delta === 'string' && delta) {
            text += delta;
            onDelta?.(text);
          }
        } catch {
          // 忽略无法解析的流片段
        }
      });
    }
    return { proofId, payload: { text }, stream: true };
  }
  try {
    const res = await API.post(url, { group, values });
    return {
      proofId: res.headers?.['x-model-trial-proof'] || '',
      payload: res.data,
      stream: false,
    };
  } catch (error) {
    const message =
      error?.response?.data?.error?.message ||
      error?.response?.data?.message ||
      error.message;
    throw new Error(message || '调用失败');
  }
}

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export async function runTrialFlow({ model, apiKind, group, values, nodes, onUpdate }) {
  const stream =
    (apiKind === 'chat_completions' || apiKind === 'responses') &&
    readStreamFlag(nodes, values);
  const submitted = await submitModelTrial({
    model,
    group,
    values,
    stream,
    onDelta: (text) => onUpdate?.({ streamingText: text, result: null, error: '' }),
  });
  if (stream) {
    onUpdate?.({
      streamingText: submitted.payload.text,
      result: submitted.payload,
      error: '',
      proofId: submitted.proofId,
      success: true,
    });
    return submitted;
  }
  const taskId = ASYNC_KINDS.includes(apiKind) ? extractTaskId(submitted.payload) : '';
  if (ASYNC_KINDS.includes(apiKind) && !taskId) {
    throw new Error('没有拿到任务编号');
  }
  if (!taskId) {
    onUpdate?.({
      result: submitted.payload,
      streamingText: '',
      error: '',
      proofId: submitted.proofId,
      success: true,
    });
    return submitted;
  }
  onUpdate?.({
    result: submitted.payload,
    streamingText: '',
    error: '',
    proofId: submitted.proofId,
    success: false,
  });
  for (let i = 0; i < 80; i += 1) {
    await sleep(3000);
    const task = await pollTrialTask(apiKind, taskId, submitted.proofId);
    const success = task.status === 'SUCCESS';
    const failed = task.status === 'FAILURE';
    onUpdate?.({
      result: task,
      streamingText: '',
      error: failed ? task.fail_reason || '任务失败' : '',
      proofId: submitted.proofId,
      success,
    });
    if (success || failed) return submitted;
  }
  throw new Error('任务等待超时');
}

export async function pollTrialTask(kind, taskId, proofId) {
  const res = await API.get(
    `/pg/trial/tasks/${encodeURIComponent(kind)}/${encodeURIComponent(taskId)}`,
    { params: { proof_id: proofId } },
  );
  if (!res.data?.success) {
    throw new Error(res.data?.message || '查询任务失败');
  }
  return res.data.data;
}
