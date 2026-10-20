import React, { useEffect, useRef } from 'react';
import {
  Input,
  TextArea,
  Switch,
  Select,
  RadioGroup,
  Radio,
  Typography,
  Button,
} from '@douyinfe/semi-ui';

const { Text } = Typography;

function Help({ text }) {
  if (!text) return null;
  return (
    <Text type='tertiary' size='small' style={{ display: 'block', marginTop: 4 }}>
      {text}
    </Text>
  );
}

function sliderBounds(node) {
  const min = Number(node.min);
  const max = Number(node.max);
  return {
    min: Number.isFinite(min) ? min : 0,
    max: Number.isFinite(max) ? max : 0,
    integer: Number.isInteger(min) && Number.isInteger(max),
  };
}

function sliderValue(value, min) {
  const parsed = value == null || value === '' ? min : Number(value);
  return Number.isFinite(parsed) ? parsed : min;
}

function formatSlider(value, integer) {
  if (!Number.isFinite(value)) return '';
  if (integer) return String(Math.round(value));
  return String(Math.round(value * 100) / 100);
}

function paintSlider(input, label, min, max, integer) {
  const numeric = Number(input.value);
  const span = max - min;
  const percent = span > 0 ? Math.min(100, Math.max(0, ((numeric - min) / span) * 100)) : 0;
  input.style.setProperty('--trial-slider-fill', `${percent}%`);
  if (label) label.textContent = formatSlider(numeric, integer);
}

function SliderField({ node, value, onChange }) {
  const { min, max, integer } = sliderBounds(node);
  const current = sliderValue(value, min);
  const span = max - min;
  const initialFill = span > 0 ? Math.min(100, Math.max(0, ((current - min) / span) * 100)) : 0;
  const inputRef = useRef(null);
  const labelRef = useRef(null);
  const dragging = useRef(false);

  useEffect(() => {
    const input = inputRef.current;
    if (!input || dragging.current) return;
    input.value = String(current);
    paintSlider(input, labelRef.current, min, max, integer);
  }, [current, min, max, integer]);

  const commit = (raw) => {
    let numeric = Number(raw);
    if (!Number.isFinite(numeric)) return;
    numeric = Math.min(max, Math.max(min, numeric));
    numeric = integer ? Math.round(numeric) : Math.round(numeric * 100) / 100;
    onChange(numeric);
  };

  return (
    <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
      <style>{`
        .trial-slider {
          -webkit-appearance: none;
          appearance: none;
          width: 100%;
          height: 20px;
          margin: 0;
          background: transparent;
          cursor: pointer;
          --trial-slider-fill: 0%;
        }
        .trial-slider::-webkit-slider-runnable-track {
          height: 4px;
          border-radius: 999px;
          background: linear-gradient(to right, var(--semi-color-primary) var(--trial-slider-fill), var(--semi-color-fill-2) var(--trial-slider-fill));
        }
        .trial-slider::-webkit-slider-thumb {
          -webkit-appearance: none;
          appearance: none;
          width: 18px;
          height: 18px;
          margin-top: -7px;
          border: none;
          border-radius: 50%;
          background: var(--semi-color-bg-0, #fff);
          box-shadow: 0 1px 4px rgba(0, 0, 0, 0.18);
          cursor: grab;
        }
        .trial-slider::-moz-range-track {
          height: 4px;
          border: none;
          border-radius: 999px;
          background: linear-gradient(to right, var(--semi-color-primary) var(--trial-slider-fill), var(--semi-color-fill-2) var(--trial-slider-fill));
        }
        .trial-slider::-moz-range-thumb {
          width: 18px;
          height: 18px;
          border: none;
          border-radius: 50%;
          background: var(--semi-color-bg-0, #fff);
          box-shadow: 0 1px 4px rgba(0, 0, 0, 0.18);
          cursor: grab;
        }
      `}</style>
      <input
        ref={inputRef}
        className='trial-slider'
        type='range'
        min={min}
        max={max}
        step='any'
        defaultValue={current}
        aria-label={node.label || '拖动条'}
        style={{ flex: 1, minWidth: 0, '--trial-slider-fill': `${initialFill}%` }}
        onPointerDown={() => {
          dragging.current = true;
          const finish = () => {
            dragging.current = false;
            if (inputRef.current) commit(inputRef.current.value);
            window.removeEventListener('pointerup', finish);
            window.removeEventListener('pointercancel', finish);
          };
          window.addEventListener('pointerup', finish);
          window.addEventListener('pointercancel', finish);
        }}
        onInput={(event) => paintSlider(event.currentTarget, labelRef.current, min, max, integer)}
        onKeyUp={(event) => commit(event.currentTarget.value)}
        onBlur={(event) => {
          dragging.current = false;
          commit(event.currentTarget.value);
        }}
      />
      <span ref={labelRef} style={{ width: 36, textAlign: 'right', flex: 'none', fontVariantNumeric: 'tabular-nums' }}>
        {formatSlider(current, integer)}
      </span>
    </div>
  );
}

function FieldControl({ node, value, onChange }) {
  if (node.widget === 'textarea' || node.widget === 'input') {
    return (
      <TextArea
        value={value || ''}
        autosize
        onChange={(next) => onChange(next)}
        placeholder={node.help || node.label}
      />
    );
  }
  if (node.widget === 'boolean') {
    return <Switch checked={!!value} onChange={(checked) => onChange(checked)} />;
  }
  if (node.widget === 'slider') {
    return <SliderField node={node} value={value} onChange={onChange} />;
  }
  if (node.widget === 'number' || node.widget === 'integer') {
    return (
      <Input
        type='number'
        value={value ?? ''}
        onChange={(next) => onChange(next === '' ? undefined : Number(next))}
      />
    );
  }
  if (node.widget === 'image_urls') {
    const text = Array.isArray(value) ? value.join('\n') : '';
    return (
      <TextArea
        value={text}
        autosize
        placeholder='每行一个地址'
        onChange={(next) =>
          onChange(
            next
              .split('\n')
              .map((item) => item.trim())
              .filter(Boolean),
          )
        }
      />
    );
  }
  if (node.widget === 'messages') {
    const messages = Array.isArray(value) && value.length ? value : [{ role: 'user', content: '' }];
    return (
      <div>
        {messages.map((message, index) => (
          <div key={index} style={{ display: 'flex', gap: 8, marginBottom: 8 }}>
            <Select
              value={message.role || 'user'}
              style={{ width: 110 }}
              optionList={[
                { value: 'system', label: 'system' },
                { value: 'user', label: 'user' },
                { value: 'assistant', label: 'assistant' },
              ]}
              onChange={(role) => {
                const next = messages.map((item, itemIndex) =>
                  itemIndex === index ? { ...item, role } : item,
                );
                onChange(next);
              }}
            />
            <TextArea
              value={message.content || ''}
              autosize
              onChange={(content) => {
                const next = messages.map((item, itemIndex) =>
                  itemIndex === index ? { ...item, content } : item,
                );
                onChange(next);
              }}
            />
          </div>
        ))}
        <Button
          size='small'
          onClick={() => onChange([...messages, { role: 'user', content: '' }])}
        >
          添加消息
        </Button>
      </div>
    );
  }
  if (node.widget === 'json') {
    const text = typeof value === 'string' ? value : value ? JSON.stringify(value, null, 2) : '';
    return (
      <TextArea
        value={text}
        autosize
        onChange={(next) => {
          try {
            onChange(JSON.parse(next));
          } catch {
            onChange(next);
          }
        }}
      />
    );
  }
  return (
    <Input
      value={value || ''}
      onChange={(next) => onChange(next)}
      placeholder={node.help || node.label}
    />
  );
}

function NodeBlock({ node, values, onChange }) {
  const value = values[node.id];
  const setValue = (next) => onChange({ ...values, [node.id]: next });

  if (node.type === 'field') {
    return (
      <div style={{ marginBottom: 24 }}>
        <Text strong>
          {node.label}
          {node.required ? ' *' : ''}
        </Text>
        <div style={{ marginTop: 10 }}>
          <FieldControl node={node} value={value} onChange={setValue} />
        </div>
        <Help text={node.help} />
      </div>
    );
  }

  if (node.type === 'object') {
    return (
      <div style={{ marginBottom: 24, padding: '4px 0 4px 16px', borderLeft: '2px solid var(--semi-color-border)' }}>
        <Text strong>{node.label}</Text>
        <Help text={node.help} />
        <TrialNodes nodes={node.nodes} values={values} onChange={onChange} />
      </div>
    );
  }

  if (node.type === 'select') {
    const custom = node.allow_custom && value && !(node.options || []).some((item) => item.value === value);
    const selected = (node.options || []).find((item) => item.value === value);
    return (
      <div style={{ marginBottom: 24 }}>
        <Text strong>
          {node.label}
          {node.required ? ' *' : ''}
        </Text>
        <div style={{ marginTop: 10 }}>
          <Select
            value={custom ? undefined : value}
            placeholder='请选择'
            style={{ width: '100%' }}
            optionList={(node.options || []).map((item) => ({ value: item.value, label: item.label }))}
            onChange={setValue}
          />
          {node.allow_custom && (
            <Input
              style={{ marginTop: 8 }}
              value={custom ? value : ''}
              placeholder='自定义值'
              onChange={setValue}
            />
          )}
        </div>
        <Help text={node.help} />
        {selected && <TrialNodes nodes={selected.nodes} values={values} onChange={onChange} />}
      </div>
    );
  }

  if (node.type === 'branch') {
    const selected = (node.options || []).find((item) => item.value === value);
    return (
      <div style={{ marginBottom: 24 }}>
        <Text strong>
          {node.label}
          {node.required ? ' *' : ''}
        </Text>
        <Help text={node.help} />
        <RadioGroup
          type='button'
          value={value}
          style={{ marginTop: 8 }}
          onChange={(event) => setValue(event.target.value)}
        >
          {(node.options || []).map((item) => (
            <Radio key={item.value} value={item.value}>
              {item.label}
            </Radio>
          ))}
        </RadioGroup>
        {selected?.help && <Help text={selected.help} />}
        {selected && <TrialNodes nodes={selected.nodes} values={values} onChange={onChange} />}
      </div>
    );
  }

  return null;
}

export function TrialNodes({ nodes, values, onChange }) {
  return (
    <div>
      {(nodes || []).map((node) => (
        <NodeBlock key={node.id} node={node} values={values} onChange={onChange} />
      ))}
    </div>
  );
}

export default function TrialForm({ nodes, values, onChange }) {
  return <TrialNodes nodes={nodes} values={values || {}} onChange={onChange} />;
}
