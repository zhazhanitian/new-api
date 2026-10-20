import React from 'react';
import { Button, Toast, Typography } from '@douyinfe/semi-ui';
import { IconCopy } from '@douyinfe/semi-icons';
import hljs from 'highlight.js/lib/core';
import jsonLang from 'highlight.js/lib/languages/json';
import { copy } from '../../helpers';

hljs.registerLanguage('json', jsonLang);

const { Text } = Typography;

function prepare(value, text) {
  const raw = text != null ? String(text) : JSON.stringify(value, null, 2);
  const trimmed = raw.trim();
  const looksJson =
    (trimmed.startsWith('{') && trimmed.endsWith('}')) ||
    (trimmed.startsWith('[') && trimmed.endsWith(']'));
  if (looksJson) {
    try {
      const pretty = JSON.stringify(JSON.parse(trimmed), null, 2);
      return {
        label: 'JSON',
        text: pretty,
        html: hljs.highlight(pretty, { language: 'json' }).value,
      };
    } catch {
      // 不是合法 JSON 时按原文展示
    }
  }
  return { label: 'TEXT', text: raw, html: null };
}

export default function JsonCodeBlock({ value, text }) {
  const { label, text: display, html } = prepare(value, text);
  const copyText = () => {
    copy(display).then((ok) => {
      if (ok) Toast.success('已复制');
      else Toast.error('复制失败');
    });
  };
  return (
    <div
      className='json-code-block'
      style={{
        border: '1px solid var(--semi-color-border)',
        borderRadius: 8,
        overflow: 'hidden',
        background: 'var(--semi-color-fill-0)',
      }}
    >
      <style>{`
        .json-code-block pre { margin: 0; }
        .json-code-block pre code {
          display: block;
          margin: 0;
          padding: 12px 14px;
          background: transparent;
          color: var(--semi-color-text-0);
          font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
          font-size: 12px;
          line-height: 1.6;
          white-space: pre-wrap;
          word-break: break-word;
        }
        .json-code-block .hljs-attr { color: var(--semi-color-primary); }
        .json-code-block .hljs-string { color: var(--semi-color-success); }
        .json-code-block .hljs-number { color: var(--semi-color-warning); }
        .json-code-block .hljs-literal { color: var(--semi-color-warning); }
      `}</style>
      <div
        style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
          padding: '4px 8px 4px 12px',
          borderBottom: '1px solid var(--semi-color-border)',
          background: 'var(--semi-color-bg-1)',
        }}
      >
        <Text type='tertiary' size='small'>{label}</Text>
        <Button size='small' theme='borderless' icon={<IconCopy />} onClick={copyText} />
      </div>
      <pre>
        {html ? (
          <code className='hljs language-json' dangerouslySetInnerHTML={{ __html: html }} />
        ) : (
          <code>{display}</code>
        )}
      </pre>
    </div>
  );
}
