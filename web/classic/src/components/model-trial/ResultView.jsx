import React, { useState } from 'react';
import { ImagePreview, Typography } from '@douyinfe/semi-ui';
import JsonCodeBlock from '../common/JsonCodeBlock';

const { Text, Paragraph } = Typography;

function collectUrls(value, found = []) {
  if (typeof value === 'string' && /^https?:\/\//.test(value)) found.push(value);
  else if (Array.isArray(value)) value.forEach((item) => collectUrls(item, found));
  else if (value && typeof value === 'object') Object.values(value).forEach((item) => collectUrls(item, found));
  return found;
}

function mediaKey(url) {
  try {
    const parsed = new URL(url);
    return `${parsed.origin}${parsed.pathname}`;
  } catch {
    return url;
  }
}

function uniqueUrls(urls) {
  const seen = new Set();
  return urls.filter((url) => {
    const key = mediaKey(url);
    if (seen.has(key)) return false;
    seen.add(key);
    return true;
  });
}

function isImageUrl(url) {
  return /\.(png|jpe?g|gif|webp|bmp)$/.test(url.split('?')[0].toLowerCase());
}

function Media({ url, onPreview }) {
  const lower = url.split('?')[0].toLowerCase();
  if (isImageUrl(url)) {
    return (
      <img
        src={url}
        alt=''
        onClick={() => onPreview(url)}
        style={{
          display: 'block',
          width: '100%',
          height: 'auto',
          maxHeight: 420,
          objectFit: 'contain',
          borderRadius: 8,
          cursor: 'zoom-in',
        }}
      />
    );
  }
  if (/\.(mp4|webm|mov)$/.test(lower)) {
    return <video src={url} controls style={{ display: 'block', width: '100%', maxHeight: 420, borderRadius: 8 }} />;
  }
  if (/\.(mp3|wav|ogg|m4a)$/.test(lower)) {
    return <audio src={url} controls style={{ width: '100%' }} />;
  }
  return (
    <a href={url} target='_blank' rel='noreferrer' style={{ wordBreak: 'break-all' }}>
      {url}
    </a>
  );
}

export default function ResultView({ result, error, streamingText }) {
  const [previewIndex, setPreviewIndex] = useState(-1);
  if (error) {
    return <Text type='danger'>{error}</Text>;
  }
  if (streamingText) {
    return <Paragraph style={{ whiteSpace: 'pre-wrap', margin: 0 }}>{streamingText}</Paragraph>;
  }
  if (!result) {
    return <Text type='tertiary'>结果会显示在这里</Text>;
  }
  if (typeof result.text === 'string') {
    return <Paragraph style={{ whiteSpace: 'pre-wrap', margin: 0 }}>{result.text}</Paragraph>;
  }
  const urls = uniqueUrls(collectUrls(result));
  const imageUrls = urls.filter(isImageUrl);
  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 12 }}>
      {urls.length > 0 && (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start', gap: 12 }}>
          {urls.map((url) => (
            <Media
              key={mediaKey(url)}
              url={url}
              onPreview={(target) => setPreviewIndex(imageUrls.indexOf(target))}
            />
          ))}
        </div>
      )}
      {imageUrls.length > 0 && (
        <ImagePreview
          src={imageUrls}
          currentIndex={previewIndex < 0 ? 0 : previewIndex}
          visible={previewIndex >= 0}
          onChange={setPreviewIndex}
          onVisibleChange={(visible) => {
            if (!visible) setPreviewIndex(-1);
          }}
        />
      )}
      <JsonCodeBlock value={result} />
    </div>
  );
}
