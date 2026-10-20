import React, { createContext, useContext, useState } from 'react';
import { Button, Input, Select, Switch, Typography } from '@douyinfe/semi-ui';
import { GripVertical } from 'lucide-react';
import { WIDGETS, blankNode, newNodeId } from '../../helpers/modelTrial';

const { Text } = Typography;

const DragContext = createContext(null);

const TYPE_LABEL = {
  field: '字段',
  select: '下拉',
  branch: '分支',
  object: '对象',
};

const cardStyle = {
  border: '1px solid var(--semi-color-border)',
  borderRadius: 12,
  padding: 16,
  marginBottom: 20,
  background: 'var(--semi-color-bg-1)',
};

const rowStyle = {
  display: 'grid',
  gridTemplateColumns: '28px 52px minmax(0, 1fr) minmax(0, 1.15fr) 64px',
  columnGap: 12,
  rowGap: 12,
  alignItems: 'center',
};

const nestedStyle = {
  marginTop: 16,
  marginLeft: 40,
  padding: 16,
  borderRadius: 10,
  background: 'var(--semi-color-fill-0)',
};

function reorder(list, from, to) {
  if (from === to || from < 0 || to < 0 || from >= list.length || to >= list.length) return list;
  const next = list.slice();
  const [item] = next.splice(from, 1);
  next.splice(to, 0, item);
  return next;
}

function updateAt(nodes, index, next) {
  return nodes.map((node, nodeIndex) => (nodeIndex === index ? next : node));
}

function NodeCard({ node, index, listId, onChange, onDelete, apiKind }) {
  const drag = useContext(DragContext);
  const patch = (partial) => onChange({ ...node, ...partial });
  const widgets = WIDGETS.filter((item) => {
    if (item.value === 'messages') return apiKind === 'chat_completions';
    if (item.value === 'input') return apiKind === 'responses';
    if (item.value === 'select') return false;
    return true;
  });
  const pathVisible = node.submit !== false;
  const dragging = drag.drag?.listId === listId && drag.drag?.index === index;
  const over = drag.over?.listId === listId && drag.over?.index === index && !dragging;

  return (
    <div
      style={{
        ...cardStyle,
        opacity: dragging ? 0.45 : 1,
        boxShadow: over ? 'inset 0 2px 0 var(--semi-color-primary)' : undefined,
      }}
      onDragOver={(event) => {
        if (drag.drag?.listId !== listId) return;
        event.preventDefault();
        event.stopPropagation();
        if (drag.over?.listId !== listId || drag.over?.index !== index) {
          drag.setOver({ listId, index });
        }
      }}
      onDrop={(event) => {
        if (drag.drag?.listId !== listId) return;
        event.preventDefault();
        event.stopPropagation();
        drag.onDrop(listId, drag.drag.index, index);
      }}
    >
      <div style={rowStyle}>
        <span
          draggable
          title='拖动排序'
          onDragStart={(event) => {
            event.stopPropagation();
            drag.setDrag({ listId, index });
            event.dataTransfer.effectAllowed = 'move';
            event.dataTransfer.setData('text/plain', String(index));
          }}
          onDragEnd={() => drag.clear()}
          style={{ cursor: 'grab', display: 'flex', alignItems: 'center', color: 'var(--semi-color-text-2)' }}
        >
          <GripVertical size={16} />
        </span>
        <Text strong style={{ lineHeight: '32px' }}>{TYPE_LABEL[node.type] || '字段'}</Text>
        <Input
          prefix='标题'
          value={node.label || ''}
          onChange={(label) => patch({ label })}
        />
        {pathVisible ? (
          <Input
            prefix='路径'
            value={node.key || ''}
            placeholder='prompt 或 metadata.ratio'
            onChange={(key) => patch({ key })}
          />
        ) : (
          <span />
        )}
        <Button
          size='small'
          type='danger'
          theme='borderless'
          style={{ height: 32 }}
          onClick={onDelete}
        >
          删除
        </Button>
        <Input
          prefix='说明'
          style={{ gridColumn: '3 / 5' }}
          value={node.help || ''}
          placeholder='告诉用户这个值怎么填'
          onChange={(help) => patch({ help })}
        />
        <div
          style={{
            gridColumn: '3 / -1',
            display: 'grid',
            gridTemplateColumns: '140px 84px 84px minmax(0, 1fr)',
            columnGap: 16,
            alignItems: 'center',
          }}
        >
          {node.type === 'field' ? (
            <Select
              value={node.widget || 'text'}
              style={{ width: 140 }}
              optionList={widgets}
              onChange={(widget) => patch({ widget })}
            />
          ) : (
            <span />
          )}
          {node.type !== 'object' ? (
            <span style={{ display: 'inline-flex', alignItems: 'center', gap: 8, height: 32 }}>
              <span>必填</span>
              <Switch size='small' checked={!!node.required} onChange={(required) => patch({ required })} />
            </span>
          ) : (
            <span />
          )}
          <span style={{ display: 'inline-flex', alignItems: 'center', gap: 8, height: 32 }}>
            <span>提交</span>
            <Switch
              size='small'
              checked={node.submit !== false}
              onChange={(submit) => patch({ submit, key: submit ? node.key : '' })}
            />
          </span>
          <span style={{ display: 'inline-flex', alignItems: 'center', gap: 12, minHeight: 32 }}>
            {node.type === 'select' && (
              <>
                <span>允许自定义</span>
                <Switch size='small' checked={!!node.allow_custom} onChange={(allow_custom) => patch({ allow_custom })} />
              </>
            )}
            {node.type === 'field' &&
              (node.widget === 'slider' || node.widget === 'number' || node.widget === 'integer') && (
                <>
                  <Input
                    prefix='最小'
                    style={{ width: 120 }}
                    value={node.min ?? ''}
                    onChange={(next) => patch({ min: next === '' ? undefined : Number(next) })}
                  />
                  <Input
                    prefix='最大'
                    style={{ width: 120 }}
                    value={node.max ?? ''}
                    onChange={(next) => patch({ max: next === '' ? undefined : Number(next) })}
                  />
                </>
              )}
          </span>
        </div>
      </div>

      {(node.type === 'select' || node.type === 'branch') && (
        <div style={nestedStyle}>
          <Text type='tertiary' size='small'>选项。每个选项下面的参数只在选中该项时出现。</Text>
          {(node.options || []).map((option, optionIndex) => (
            <div
              key={`${node.id}-option-${optionIndex}`}
              style={{
                marginTop: 16,
                padding: 16,
                borderRadius: 10,
                background: 'var(--semi-color-bg-1)',
                border: '1px solid var(--semi-color-border)',
              }}
            >
              <div style={{ display: 'grid', gridTemplateColumns: 'minmax(0, 1fr) minmax(0, 1fr) 64px', gap: 12, alignItems: 'center' }}>
                <Input
                  prefix='选项值'
                  value={option.value || ''}
                  onChange={(value) => {
                    const options = node.options.map((item, itemIndex) =>
                      itemIndex === optionIndex ? { ...item, value } : item,
                    );
                    patch({ options });
                  }}
                />
                <Input
                  prefix='选项标题'
                  value={option.label || ''}
                  onChange={(label) => {
                    const options = node.options.map((item, itemIndex) =>
                      itemIndex === optionIndex ? { ...item, label } : item,
                    );
                    patch({ options });
                  }}
                />
                <Button
                  size='small'
                  type='danger'
                  theme='borderless'
                  onClick={() => patch({ options: node.options.filter((_, itemIndex) => itemIndex !== optionIndex) })}
                >
                  删除
                </Button>
                <Input
                  prefix='选项说明'
                  style={{ gridColumn: '1 / -1' }}
                  value={option.help || ''}
                  onChange={(help) => {
                    const options = node.options.map((item, itemIndex) =>
                      itemIndex === optionIndex ? { ...item, help } : item,
                    );
                    patch({ options });
                  }}
                />
              </div>
              <div style={{ marginTop: 16 }}>
                <Text type='tertiary' size='small'>选中「{option.label || '此选项'}」后显示的参数</Text>
                <NodeList
                  listId={`${listId}/${node.id}/option/${optionIndex}`}
                  apiKind={apiKind}
                  nodes={option.nodes || []}
                  onChange={(nodes) => {
                    const options = node.options.map((item, itemIndex) =>
                      itemIndex === optionIndex ? { ...item, nodes } : item,
                    );
                    patch({ options });
                  }}
                />
              </div>
            </div>
          ))}
          <Button
            style={{ marginTop: 16 }}
            onClick={() =>
              patch({
                options: [...(node.options || []), { value: '', label: '新选项', help: '', nodes: [] }],
              })
            }
          >
            添加选项
          </Button>
        </div>
      )}

      {node.type === 'object' && (
        <div style={nestedStyle}>
          <Text type='tertiary' size='small'>对象里的参数。相对路径会接到这个对象的路径后面。</Text>
          <NodeList
            listId={`${listId}/${node.id}/object`}
            apiKind={apiKind}
            nodes={node.nodes || []}
            onChange={(nodes) => patch({ nodes })}
          />
        </div>
      )}
    </div>
  );
}

export function NodeList({ nodes, onChange, apiKind, listId = 'root' }) {
  const drag = useContext(DragContext);
  return (
    <div
      style={{ marginTop: 16 }}
      onDragOver={(event) => {
        if (drag?.drag?.listId === listId) event.preventDefault();
      }}
    >
      {(nodes || []).map((node, index) => (
        <NodeCard
          key={node.id || index}
          listId={listId}
          index={index}
          apiKind={apiKind}
          node={node}
          onChange={(next) => onChange(updateAt(nodes, index, next))}
          onDelete={() => onChange(nodes.filter((_, nodeIndex) => nodeIndex !== index))}
        />
      ))}
      <div style={{ display: 'flex', gap: 12, marginTop: 4 }}>
        {[
          ['field', '加字段'],
          ['select', '加下拉'],
          ['branch', '加分支'],
          ['object', '加对象'],
        ].map(([type, label]) => (
          <Button key={type} onClick={() => onChange([...(nodes || []), blankNode(type)])}>
            {label}
          </Button>
        ))}
      </div>
    </div>
  );
}

export function ensureNodeIds(nodes) {
  return (nodes || []).map((node) => ({
    ...node,
    id: node.id || newNodeId(),
    nodes: node.nodes ? ensureNodeIds(node.nodes) : node.nodes,
    options: node.options
      ? node.options.map((option) => ({
          ...option,
          nodes: ensureNodeIds(option.nodes),
        }))
      : node.options,
  }));
}

export default function NodeEditor({ nodes, onChange, apiKind }) {
  const [drag, setDrag] = useState(null);
  const [over, setOver] = useState(null);
  const clear = () => {
    setDrag(null);
    setOver(null);
  };
  const onDrop = (listId, from, to) => {
    const current = nodes || [];
    onChange(listId === 'root' ? reorder(current, from, to) : reorderInTree(current, listId, from, to));
    clear();
  };

  return (
    <DragContext.Provider value={{ drag, over, setDrag, setOver, clear, onDrop }}>
      <NodeList listId='root' apiKind={apiKind} nodes={nodes} onChange={onChange} />
    </DragContext.Provider>
  );
}

function reorderInTree(nodes, listId, from, to) {
  const parts = listId.split('/').slice(1);
  return walk(nodes, parts);

  function walk(list, rest) {
    if (rest.length === 0) return reorder(list, from, to);
    const nodeId = rest[0];
    const kind = rest[1];
    return list.map((node) => {
      if (node.id !== nodeId) return node;
      if (kind === 'object') {
        return { ...node, nodes: walk(node.nodes || [], rest.slice(2)) };
      }
      const index = Number(rest[2]);
      return {
        ...node,
        options: (node.options || []).map((option, optionIdx) =>
          optionIdx === index ? { ...option, nodes: walk(option.nodes || [], rest.slice(3)) } : option,
        ),
      };
    });
  }
}
