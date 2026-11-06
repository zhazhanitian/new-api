import React from 'react';

function urlLines(url) {
  if (!url) return [];
  const queryAt = url.indexOf('?');
  if (queryAt < 0) return [url];
  const base = url.slice(0, queryAt);
  const params = url
    .slice(queryAt + 1)
    .split('&')
    .filter(Boolean);
  return [base, ...params];
}

const ResultUrlBlock = ({ url, compact = false }) => {
  const lines = urlLines(url);
  return (
    <div
      style={{
        marginTop: compact ? 0 : 16,
        padding: '10px 12px',
        borderRadius: 6,
        background: 'var(--semi-color-fill-0)',
        display: 'flex',
        flexDirection: 'column',
        gap: 6,
      }}
    >
      {lines.map((line) => (
        <div
          key={line}
          style={{
            overflowX: 'auto',
            whiteSpace: 'nowrap',
            fontFamily: 'ui-monospace, SFMono-Regular, Menlo, Consolas, monospace',
            fontSize: 12,
            lineHeight: '20px',
            color: 'var(--semi-color-text-2)',
          }}
        >
          {line}
        </div>
      ))}
    </div>
  );
};

export default ResultUrlBlock;
