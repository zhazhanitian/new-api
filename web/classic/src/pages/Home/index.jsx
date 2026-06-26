/*
Copyright (C) 2025 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/

import React, {
  useCallback,
  useContext,
  useEffect,
  useRef,
  useState,
} from 'react';
import {
  Button,
  Typography,
  Input,
  ScrollList,
  ScrollItem,
} from '@douyinfe/semi-ui';
import {
  API,
  showError,
  copy,
  showSuccess,
  getSystemName,
} from '../../helpers';
import { useIsMobile } from '../../hooks/common/useIsMobile';
import { API_ENDPOINTS } from '../../constants/common.constant';
import { StatusContext } from '../../context/Status';
import { useActualTheme } from '../../context/Theme';
import { marked } from 'marked';
import { useTranslation } from 'react-i18next';
import {
  IconCopy,
  IconPlay,
  IconFile,
  IconArrowRight,
} from '@douyinfe/semi-icons';
import { Link } from 'react-router-dom';
import NoticeModal from '../../components/layout/NoticeModal';
import {
  Moonshot, OpenAI, XAI, Zhipu, Volcengine, Cohere, Claude, Gemini,
  Suno, Minimax, Wenxin, Spark, Qingyan, DeepSeek, Qwen, Midjourney,
  Grok, AzureAI, Hunyuan, Xinference,
} from '@lobehub/icons';
import capabilityLlmImg from '../../assets/capability-llm-placeholder.svg';
import capabilityImageImg from '../../assets/capability-image-placeholder.svg';
import capabilityVideoImg from '../../assets/capability-video-placeholder.svg';
import capabilityAudioImg from '../../assets/capability-audio-placeholder.svg';

const { Text } = Typography;

/* ─────────────────────────────────────────────────────
   滚动触发渐入
───────────────────────────────────────────────────── */
function Reveal({ children, delay = 0, className = '' }) {
  const ref = useRef(null);
  useEffect(() => {
    const el = ref.current;
    if (!el) return;
    const obs = new IntersectionObserver(
      ([e]) => {
        if (e.isIntersecting) {
          setTimeout(() => el.classList.add('home-revealed'), delay);
          obs.disconnect();
        }
      },
      { threshold: 0.06, rootMargin: '0px 0px -20px 0px' },
    );
    obs.observe(el);
    return () => obs.disconnect();
  }, [delay]);
  return <div ref={ref} className={`home-reveal ${className}`}>{children}</div>;
}

/* ─────────────────────────────────────────────────────
   数字滚动卡片
───────────────────────────────────────────────────── */
function StatCard({ value, label, detail, delay = 0 }) {
  const [display, setDisplay] = useState('');
  const done = useRef(false);
  const ref = useRef(null);

  const run = useCallback(() => {
    if (done.current) return;
    done.current = true;
    const m = value.match(/([\d.]+)/);
    if (!m) { setDisplay(value); return; }
    const num = m[1], end = parseFloat(num);
    const pre = value.slice(0, value.indexOf(num));
    const suf = value.slice(value.indexOf(num) + num.length);
    const dec = num.includes('.');
    const t0 = performance.now();
    const tick = (now) => {
      const p = Math.min((now - t0) / 1200, 1), e = 1 - (1 - p) ** 3;
      setDisplay(pre + (dec ? (end * e).toFixed(1) : Math.round(end * e)) + suf);
      if (p < 1) requestAnimationFrame(tick); else setDisplay(value);
    };
    requestAnimationFrame(tick);
  }, [value]);

  useEffect(() => {
    const obs = new IntersectionObserver(
      ([e]) => { if (e.isIntersecting) { setTimeout(run, delay); obs.disconnect(); } },
      { threshold: 0.3 },
    );
    if (ref.current) obs.observe(ref.current);
    return () => obs.disconnect();
  }, [run, delay]);

  return (
    <div
      ref={ref}
      className='flex flex-col gap-3 rounded-2xl border border-semi-color-border bg-semi-color-bg-1 px-6 py-7 transition-all duration-300 hover:-translate-y-1 hover:shadow-lg'
      style={{ boxShadow: '0 2px 10px rgba(0,0,0,0.04)' }}
    >
      {/* 数字 — 对标 V1 的 34px，这里给到 44px */}
      <div
        className='font-extrabold leading-none tracking-tight text-semi-color-text-0'
        style={{ fontSize: 44, fontVariantNumeric: 'tabular-nums' }}
      >
        {display || value}
      </div>
      <div className='text-[15px] font-bold text-semi-color-text-0'>{label}</div>
      <div className='text-xs font-medium leading-relaxed text-semi-color-text-2'>{detail}</div>
    </div>
  );
}

/* ─────────────────────────────────────────────────────
   代码预览
───────────────────────────────────────────────────── */
function CodePreview({ serverAddress }) {
  const [tab, setTab] = useState('request');
  const [copied, setCopied] = useState(false);

  const req = `curl "${serverAddress}/v1/chat/completions" \\
  -H "Content-Type: application/json" \\
  -H "Authorization: Bearer YOUR_API_KEY" \\
  -d '{
    "model": "gpt-4o",
    "messages": [{ "role": "user", "content": "你好" }]
  }'`;

  const res = `{
  "id": "chatcmpl-abc123",
  "object": "chat.completion",
  "model": "gpt-4o",
  "choices": [{
    "message": { "role": "assistant", "content": "你好！有什么可以帮你的？" },
    "finish_reason": "stop"
  }],
  "usage": { "prompt_tokens": 9, "completion_tokens": 12, "total_tokens": 21 }
}`;

  const doCopy = async () => {
    await navigator.clipboard.writeText(tab === 'request' ? req : res);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div
      className='overflow-hidden rounded-2xl border border-semi-color-border'
      style={{ boxShadow: '0 2px 10px rgba(0,0,0,0.04)' }}
    >
      <div className='flex items-center justify-between border-b border-semi-color-border bg-semi-color-fill-0 px-5 py-2.5'>
        <div className='flex gap-1'>
          {[['request', 'Request'], ['response', 'Response']].map(([k, l]) => (
            <button
              key={k} type='button' onClick={() => setTab(k)}
              className={`rounded-lg px-3.5 py-1.5 text-[13px] font-semibold transition-colors ${
                tab === k
                  ? 'bg-semi-color-primary text-semi-color-white'
                  : 'text-semi-color-text-2 hover:bg-semi-color-fill-1 hover:text-semi-color-text-0'
              }`}
            >{l}</button>
          ))}
        </div>
        <button
          type='button' onClick={doCopy}
          className='flex items-center gap-1.5 rounded-lg px-3.5 py-1.5 text-[13px] font-semibold text-semi-color-text-2 hover:bg-semi-color-fill-1 hover:text-semi-color-text-0 transition-colors'
        >
          <IconCopy style={{ fontSize: 14 }} />{copied ? '已复制' : '复制'}
        </button>
      </div>
      <div className='overflow-x-auto bg-semi-color-bg-0 px-6 py-6'>
        <pre
          className='m-0 p-0 text-sm leading-7 text-semi-color-text-0 whitespace-pre'
          style={{ fontFamily: "'SF Mono','Fira Code',Menlo,Consolas,monospace" }}
        >
          <code>{tab === 'request' ? req : res}</code>
        </pre>
      </div>
    </div>
  );
}

/* ─────────────────────────────────────────────────────
   静态数据
───────────────────────────────────────────────────── */
const ADVANTAGES = [
  {
    num: '01',
    title: '一次接入，覆盖全部模型',
    desc: '无需为每个模型单独维护 SDK 和鉴权逻辑，统一 OpenAI 兼容接口，切换模型只需改一个参数。',
    color: 'var(--semi-color-primary)',
    bg: 'var(--semi-color-primary-light-default)',
  },
  {
    num: '02',
    title: '统一账单与调用日志',
    desc: '所有模型的 Token 消耗、费用与请求日志集中管理，不再跨多个平台核对账单。',
    color: 'var(--semi-color-success)',
    bg: 'var(--semi-color-success-light-default)',
  },
  {
    num: '03',
    title: '生产级可用性保障',
    desc: '自动故障切换、限速保护与负载均衡，让模型调用在生产环境中稳定运行。',
    color: 'var(--semi-color-info)',
    bg: 'var(--semi-color-info-light-default)',
  },
];

const CAPABILITIES = [
  {
    eyebrow: 'LLM & Assistant',
    title: 'LLM 与对话 API',
    desc: '适合问答助手、推理与结构化输出。覆盖 GPT、Claude、Gemini、DeepSeek 等主流大模型。',
    tags: ['Chat Completions', 'Responses API', 'Embeddings', 'Rerank'],
    image: capabilityLlmImg,
  },
  {
    eyebrow: 'Image Models',
    title: '图像生成 API',
    desc: '适合海报、商品图与品牌视觉生成。覆盖 DALL·E、Midjourney、Flux、即梦等。',
    tags: ['DALL·E 3', 'Midjourney', 'Flux', '即梦'],
    image: capabilityImageImg,
  },
  {
    eyebrow: 'Video Models',
    title: '视频生成 API',
    desc: '适合营销视频、演示短片与动态内容。覆盖可灵、Sora、Runway、Vidu 等。',
    tags: ['可灵', 'Runway', 'Sora', 'Vidu'],
    image: capabilityVideoImg,
  },
  {
    eyebrow: 'Audio & Speech',
    title: '音频生成 API',
    desc: '适合配音、音乐与品牌音频生成。覆盖 OpenAI Audio、Suno 等主流音频模型。',
    tags: ['TTS', 'Whisper', 'Suno', '语音克隆'],
    image: capabilityAudioImg,
  },
];

const PROVIDERS = [
  <Moonshot key='ms' size={44} />, <OpenAI key='oa' size={44} />, <XAI key='xai' size={44} />,
  <Zhipu.Color key='zp' size={44} />, <Volcengine.Color key='ve' size={44} />, <Cohere.Color key='ch' size={44} />,
  <Claude.Color key='cl' size={44} />, <Gemini.Color key='gm' size={44} />, <Suno key='su' size={44} />,
  <Minimax.Color key='mm' size={44} />, <Wenxin.Color key='wx' size={44} />, <Spark.Color key='sp' size={44} />,
  <Qingyan.Color key='qy' size={44} />, <DeepSeek.Color key='ds' size={44} />, <Qwen.Color key='qw' size={44} />,
  <Midjourney key='mj' size={44} />, <Grok key='gk' size={44} />, <AzureAI.Color key='az' size={44} />,
  <Hunyuan.Color key='hy' size={44} />, <Xinference.Color key='xi' size={44} />,
];

/* ─────────────────────────────────────────────────────
   共用 Section 标题
───────────────────────────────────────────────────── */
function SectionHead({ kicker, title, subtitle, right }) {
  return (
    <div className='mb-10 flex flex-wrap items-end justify-between gap-4'>
      <div>
        <p className='mb-2 text-[11px] font-bold uppercase tracking-[0.1em] text-semi-color-primary'>{kicker}</p>
        <h2
          className='font-bold tracking-tight text-semi-color-text-0'
          style={{ fontSize: 'clamp(22px,2.6vw,32px)' }}
        >{title}</h2>
        {subtitle && (
          <p className='mt-2 text-[15px] text-semi-color-text-1'>{subtitle}</p>
        )}
      </div>
      {right}
    </div>
  );
}

/* ─────────────────────────────────────────────────────
   主组件
───────────────────────────────────────────────────── */
const Home = () => {
  const { t, i18n } = useTranslation();
  const [statusState] = useContext(StatusContext);
  const actualTheme = useActualTheme();
  const [pageLoaded, setPageLoaded] = useState(false);
  const [pageContent, setPageContent] = useState('');
  const [noticeVisible, setNoticeVisible] = useState(false);
  const isMobile = useIsMobile();
  const docsLink = statusState?.status?.docs_link || '';
  const serverAddress =
    statusState?.status?.server_address || `${window.location.origin}`;
  const endpointItems = API_ENDPOINTS.map((e) => ({ value: e }));
  const [epIdx, setEpIdx] = useState(0);

  /* 备案：首页 title 必须含公司名称 */
  useEffect(() => {
    const sys = getSystemName() || 'MAPI';
    document.title = `南京白鲸汇智能科技有限公司 - ${sys}`;
    return () => { document.title = sys; };
  }, []);

  const loadContent = async () => {
    setPageContent(localStorage.getItem('home_page_content') || '');
    const res = await API.get('/api/home_page_content');
    const { success, message, data } = res.data;
    if (success) {
      const content = data.startsWith('https://') ? data : marked.parse(data);
      setPageContent(content);
      localStorage.setItem('home_page_content', content);
      if (data.startsWith('https://')) {
        const iframe = document.querySelector('iframe');
        if (iframe) iframe.onload = () => {
          iframe.contentWindow.postMessage({ themeMode: actualTheme }, '*');
          iframe.contentWindow.postMessage({ lang: i18n.language }, '*');
        };
      }
    } else {
      showError(message);
      setPageContent('加载失败...');
    }
    setPageLoaded(true);
  };

  const handleCopy = async () => {
    if (await copy(serverAddress)) showSuccess(t('已复制到剪切板'));
  };

  useEffect(() => {
    const timer = setInterval(
      () => setEpIdx((p) => (p + 1) % endpointItems.length),
      3000,
    );
    return () => clearInterval(timer);
  }, [endpointItems.length]);

  useEffect(() => {
    (async () => {
      const last = localStorage.getItem('notice_close_date');
      if (last !== new Date().toDateString()) {
        try {
          const r = await API.get('/api/notice');
          if (r.data.success && r.data.data?.trim()) setNoticeVisible(true);
        } catch {}
      }
    })();
  }, []);

  useEffect(() => { loadContent().then(); }, []);

  if (pageLoaded && pageContent) {
    return (
      <div className='overflow-x-hidden w-full'>
        {pageContent.startsWith('https://') ? (
          <iframe src={pageContent} className='w-full h-screen border-none' />
        ) : (
          <div className='mt-[60px]' dangerouslySetInnerHTML={{ __html: pageContent }} />
        )}
      </div>
    );
  }

  /* 最大宽度容器 — 1200px 对齐 V1 */
  const W = 'mx-auto w-full max-w-[1400px] px-6 md:px-12';

  return (
    <div className='w-full overflow-x-hidden'>
      <NoticeModal
        visible={noticeVisible}
        onClose={() => setNoticeVisible(false)}
        isMobile={isMobile}
      />

      {/* ══════════════════════════════════════
          §1  HERO
          目标：V1 约 548px 高，大字幕 74px
      ══════════════════════════════════════ */}
      <section
        className='relative overflow-hidden border-b border-semi-color-border'
        style={{ minHeight: isMobile ? 480 : 600 }}
      >
        <div className='blur-ball blur-ball-indigo' />
        <div className='blur-ball blur-ball-teal' />

        {/* 水印 */}
        <div
          aria-hidden='true'
          className='home-wordmark-float pointer-events-none absolute left-1/2 -translate-x-1/2 select-none whitespace-nowrap font-black uppercase leading-none tracking-[0.14em] text-semi-color-text-0'
          style={{ fontSize: 'clamp(100px,15vw,220px)', opacity: 0.03, top: 80 }}
        >
          MAPI
        </div>

        {/* V1 同款 hero center：宽度 min(100%,880px) */}
        <div
          className='relative z-10 mx-auto w-full max-w-[880px] px-6 text-center'
          style={{ paddingTop: isMobile ? 80 : 110, paddingBottom: isMobile ? 64 : 80 }}
        >
          {/* Badge */}
          <div className='home-anim home-anim-d1 mb-8 inline-flex items-center rounded-full border border-semi-color-primary-light-hover bg-semi-color-primary-light-default px-4 py-2 text-xs font-bold uppercase tracking-widest text-semi-color-primary'>
            MAPI · AI 模型 API 聚合平台
          </div>

          {/* 主标题 — 完全对照 V1：74px / letter-spacing -0.065em */}
          <h1
            className='home-anim home-anim-d2 w-full font-bold text-semi-color-text-0'
            style={{
              fontSize: isMobile ? '36px' : '74px',
              lineHeight: 0.96,
              letterSpacing: '-0.065em',
              maxWidth: 860,
              margin: '0 auto',
              marginBottom: 'clamp(20px, 2.5vw, 32px)',
            }}
          >
            <span style={{ display: 'block' }}>统一接入主流 AI 模型</span>
            <span style={{ display: 'block', marginTop: 6 }}>
              <span className='shine-text' style={{ color: 'var(--semi-color-primary)' }}>
                一套 API 完成接入与治理
              </span>
            </span>
          </h1>

          {/* 副标题 */}
          <p
            className='home-anim home-anim-d3 text-semi-color-text-1'
            style={{
              fontSize: 'clamp(15px, 1.4vw, 18px)',
              lineHeight: 1.8,
              maxWidth: 640,
              margin: '0 auto',
              marginBottom: 'clamp(24px, 3vw, 40px)',
            }}
          >
            覆盖 LLM、图像、视频与音频，兼容 OpenAI 协议，
            统一鉴权、日志与账单，无需改动现有代码。
          </p>

          {/* BASE URL */}
          <div
            className='home-anim home-anim-d4 w-full'
            style={{ maxWidth: 560, margin: '0 auto', marginBottom: 'clamp(20px, 2.5vw, 32px)' }}
          >
            <p className='mb-3 text-xs font-semibold uppercase tracking-wider text-semi-color-text-2'>
              多模型统一接入，只需将基址替换为：
            </p>
            <Input
              readonly
              value={serverAddress}
              className='!rounded-full'
              size={isMobile ? 'default' : 'large'}
              suffix={
                <div className='flex items-center gap-2'>
                  <ScrollList bodyHeight={32} style={{ border: 'unset', boxShadow: 'unset' }}>
                    <ScrollItem
                      mode='wheel' cycled
                      list={endpointItems}
                      selectedIndex={epIdx}
                      onSelect={({ index }) => setEpIdx(index)}
                    />
                  </ScrollList>
                  <Button
                    type='primary' icon={<IconCopy />}
                    onClick={handleCopy} className='!rounded-full'
                  />
                </div>
              }
            />
          </div>

          {/* CTA */}
          <div className='home-anim home-anim-d5 flex flex-wrap justify-center gap-3'>
            <Link to='/console'>
              <Button
                theme='solid' type='primary'
                size={isMobile ? 'default' : 'large'}
                icon={<IconPlay />} className='!rounded-3xl !px-8'
              >
                {t('获取密钥')}
              </Button>
            </Link>
            {docsLink && (
              <Button
                size={isMobile ? 'default' : 'large'}
                icon={<IconFile />} className='!rounded-3xl !px-7'
                onClick={() => window.open(docsLink, '_blank')}
              >
                {t('文档')}
              </Button>
            )}
          </div>
        </div>
      </section>

      {/* ══════════════════════════════════════
          §2  数据指标
      ══════════════════════════════════════ */}
      <section className={`${W} py-16`}>
        <div className='grid grid-cols-2 gap-4 md:grid-cols-4'>
          <StatCard value='100+' label='接入模型'      detail='LLM · 图像 · 视频 · 音频' delay={0}   />
          <StatCard value='99.9%' label='平台可用性'   detail='SLA 保障，适合生产环境'    delay={100} />
          <StatCard value='<50ms' label='转发链路延迟'  detail='低延迟，高并发'            delay={200} />
          <StatCard value='70%'   label='接入成本节省'  detail='统一 API · Key · 日志'    delay={300} />
        </div>
      </section>

      {/* ══════════════════════════════════════
          §3  Quick Start
      ══════════════════════════════════════ */}
      <section className={`${W} border-t border-semi-color-border py-16`}>
        <SectionHead
          kicker='Quick Start'
          title='一行命令，即刻调用'
          subtitle='兼容 OpenAI 接口，现有代码无需改动，只需替换 base URL。'
          right={
            docsLink && (
              <Button type='tertiary' onClick={() => window.open(docsLink, '_blank')}>
                查看完整文档
              </Button>
            )
          }
        />
        <Reveal>
          <CodePreview serverAddress={serverAddress} />
        </Reveal>
      </section>

      {/* ══════════════════════════════════════
          §4  Why MAPI — V1 风格：大主卡 + 3 小卡
      ══════════════════════════════════════ */}
      <section className={`${W} border-t border-semi-color-border py-16`}>
        <SectionHead
          kicker='Why MAPI'
          title='为什么选择 MAPI'
          subtitle='为多模型接入建立统一的生产标准。'
        />

        {/* 主 Lead 卡 */}
        <Reveal className='mb-5'>
          <div
            className='overflow-hidden rounded-2xl border border-semi-color-border bg-semi-color-bg-1'
            style={{
              boxShadow: '0 2px 10px rgba(0,0,0,0.04)',
              backgroundImage: 'radial-gradient(circle at 90% 12%, var(--semi-color-primary-light-default), transparent 28%), linear-gradient(180deg, var(--semi-color-bg-1) 0%, var(--semi-color-bg-0) 100%)',
            }}
          >
            <div className='grid gap-6 p-8 md:grid-cols-[1fr_260px] md:p-10'>
              {/* 左侧内容 */}
              <div>
                <div className='mb-4 inline-flex items-center rounded-full bg-semi-color-primary-light-default px-3 py-1 text-[11px] font-bold uppercase tracking-widest text-semi-color-primary'>
                  Platform Standard
                </div>
                <h3
                  className='mb-4 font-bold tracking-tight text-semi-color-text-0'
                  style={{ fontSize: 'clamp(18px, 2vw, 26px)', lineHeight: 1.3 }}
                >
                  让多模型接入遵循同一套生产标准
                </h3>
                <p className='mb-5 max-w-lg text-[15px] leading-relaxed text-semi-color-text-1'>
                  统一 API 承接主流模型能力，统一控制台管理 Key、日志与账单，
                  让研发接入、客户演示与正式上线使用同一套平台口径。
                </p>
                <div className='flex flex-wrap gap-2'>
                  {['统一协议', '统一鉴权', '统一计费'].map((p) => (
                    <span
                      key={p}
                      className='rounded-full border border-semi-color-border bg-semi-color-bg-0 px-3 py-1 text-xs font-semibold text-semi-color-text-1'
                    >
                      {p}
                    </span>
                  ))}
                </div>
              </div>

              {/* 右侧 Pillar 指标 */}
              <div className='flex flex-col gap-3'>
                {[
                  { label: '协议标准', value: 'OpenAI Compatible' },
                  { label: '运营治理', value: '日志 / 账单 / Key' },
                  { label: '交付表达', value: '文档 / 控制台 / API' },
                ].map((item) => (
                  <div
                    key={item.label}
                    className='flex flex-col justify-between rounded-xl border border-semi-color-border bg-semi-color-bg-0 px-4 py-3'
                    style={{ minHeight: 72 }}
                  >
                    <span className='text-[10px] font-bold uppercase tracking-widest text-semi-color-text-2'>
                      {item.label}
                    </span>
                    <strong className='mt-2 text-[15px] font-bold leading-snug text-semi-color-text-0'>
                      {item.value}
                    </strong>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </Reveal>

        {/* 3 个特性小卡 */}
        <div className='grid gap-4 md:grid-cols-3'>
          {ADVANTAGES.map((a, i) => (
            <Reveal key={a.num} delay={i * 80}>
              <div
                className='flex flex-col rounded-2xl border border-semi-color-border bg-semi-color-bg-1 p-7 transition-all duration-300 hover:-translate-y-1 hover:shadow-lg'
                style={{ boxShadow: '0 2px 10px rgba(0,0,0,0.04)', height: '100%' }}
              >
                <div
                  className='mb-5 inline-flex h-9 w-9 items-center justify-center rounded-xl text-sm font-extrabold'
                  style={{ background: a.bg, color: a.color }}
                >
                  {a.num}
                </div>
                <h3 className='mb-3 text-[17px] font-bold text-semi-color-text-0'>{a.title}</h3>
                <p className='flex-1 text-sm leading-relaxed text-semi-color-text-1'>{a.desc}</p>
              </div>
            </Reveal>
          ))}
        </div>
      </section>

      {/* ══════════════════════════════════════
          §5  AI 能力 — V1 风格：带视觉头部的卡片
      ══════════════════════════════════════ */}
      <section className={`${W} border-t border-semi-color-border py-16`}>
        <SectionHead
          kicker='Capabilities'
          title='覆盖的 AI 能力'
          subtitle='一个账号，统一接入四类 AI 能力。'
          right={
            <Link to='/pricing'>
              <Button type='tertiary'>
                查看全部模型 <IconArrowRight style={{ marginLeft: 4 }} />
              </Button>
            </Link>
          }
        />
        <div className='grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4' style={{ alignItems: 'stretch' }}>
          {CAPABILITIES.map((c, i) => (
            <Reveal key={c.title} delay={i * 60}>
              <Link to='/pricing' className='group flex h-full' style={{ height: '100%' }}>
                <div
                  className='relative flex h-full w-full flex-col overflow-hidden rounded-2xl border border-semi-color-border bg-semi-color-bg-1 p-3 transition-all duration-300 group-hover:-translate-y-1.5 group-hover:shadow-[0_18px_40px_rgba(15,23,42,0.10)]'
                  style={{ boxShadow: '0 2px 10px rgba(0,0,0,0.04)' }}
                >
                  {/* V1 同款媒体区域：4:3 SVG 插图 */}
                  <div
                    className='w-full overflow-hidden rounded-[18px] border border-semi-color-border'
                    style={{
                      aspectRatio: '4/3',
                      flexShrink: 0,
                      background: 'linear-gradient(180deg, #fbfcff 0%, #f7f9fc 100%)',
                    }}
                  >
                    <img
                      src={c.image}
                      alt={c.title}
                      className='block h-full w-full object-cover'
                    />
                  </div>

                  {/* 内容区 — flex-1 撑满剩余高度，tags 钉在底部 */}
                  <div className='flex flex-1 flex-col px-1.5 pb-2 pt-4'>
                    <div className='mb-3 inline-flex self-start items-center rounded-full bg-semi-color-primary-light-default px-2.5 py-1.5 text-[11px] font-bold uppercase tracking-wide text-semi-color-primary'>
                      {c.eyebrow}
                    </div>
                    <h3 className='mb-2 text-[17px] font-bold leading-snug text-semi-color-text-0'>{c.title}</h3>
                    <p className='mb-4 flex-1 text-[13px] leading-relaxed text-semi-color-text-1'>{c.desc}</p>
                    <div className='flex flex-wrap gap-1.5'>
                      {c.tags.map((tag) => (
                        <span
                          key={tag}
                          className='rounded-full border border-semi-color-border bg-semi-color-fill-0 px-2.5 py-0.5 text-[11px] font-semibold text-semi-color-text-2'
                        >
                          {tag}
                        </span>
                      ))}
                    </div>
                  </div>

                  {/* hover 箭头 */}
                  <div className='absolute bottom-5 right-5 flex h-8 w-8 translate-x-[-6px] items-center justify-center rounded-lg bg-semi-color-fill-0 text-semi-color-text-2 opacity-0 transition-all duration-200 group-hover:translate-x-0 group-hover:bg-semi-color-primary-light-default group-hover:text-semi-color-primary group-hover:opacity-100'>
                    <IconArrowRight style={{ fontSize: 14 }} />
                  </div>
                </div>
              </Link>
            </Reveal>
          ))}
        </div>
      </section>

      {/* ══════════════════════════════════════
          §6  模型供应商 — 无限横向滚动轮播
      ══════════════════════════════════════ */}
      <section className={`${W} border-t border-semi-color-border py-16`}>
        <SectionHead
          kicker='Supported Providers'
          title='支持众多大模型供应商'
          subtitle='持续扩展中，覆盖国内外主流模型平台。'
        />
        <Reveal>
          {/* 在内容区宽度内 overflow，两侧渐隐 */}
          <div
            className='overflow-hidden'
            style={{
              maskImage: 'linear-gradient(to right, transparent 0%, black 8%, black 92%, transparent 100%)',
              WebkitMaskImage: 'linear-gradient(to right, transparent 0%, black 8%, black 92%, transparent 100%)',
            }}
          >
            <div className='providers-track py-2'>
              {[...PROVIDERS, ...PROVIDERS].map((icon, i) => (
                <div
                  key={i}
                  className='mx-3 flex h-20 w-20 flex-shrink-0 items-center justify-center rounded-2xl border border-semi-color-border bg-semi-color-bg-1 transition-all duration-200 hover:-translate-y-1 hover:shadow-md'
                >
                  {icon}
                </div>
              ))}
            </div>
          </div>
        </Reveal>
      </section>

      {/* ══════════════════════════════════════
          §7  Bottom CTA
      ══════════════════════════════════════ */}
      <section className={`${W} border-t border-semi-color-border py-16 pb-24`}>
        <Reveal>
          <div
            className='flex flex-col items-start justify-between gap-8 overflow-hidden rounded-2xl border border-semi-color-border bg-semi-color-bg-1 px-10 py-12 md:flex-row md:items-center'
            style={{
              backgroundImage:
                'radial-gradient(circle at 94% 10%, var(--semi-color-primary-light-default), transparent 40%), radial-gradient(circle at 6% 90%, var(--semi-color-info-light-default), transparent 36%)',
              boxShadow: '0 4px 24px rgba(0,0,0,0.06)',
            }}
          >
            <div style={{ maxWidth: 560 }}>
              <p className='mb-2.5 text-[11px] font-bold uppercase tracking-[0.1em] text-semi-color-primary'>
                Get Started
              </p>
              <h2
                className='mb-3 font-extrabold tracking-tight text-semi-color-text-0'
                style={{ fontSize: 'clamp(22px, 2.6vw, 32px)' }}
              >
                开始接入 MAPI，今天就能上线
              </h2>
              <p className='text-[15px] leading-relaxed text-semi-color-text-1'>
                注册即可获取 API Key，接入 100+ 模型，统一管理调用日志与账单。
              </p>
            </div>
            <div className='flex flex-shrink-0 flex-wrap gap-3'>
              <Link to='/console'>
                <Button
                  theme='solid' type='primary' size='large'
                  icon={<IconPlay />} className='!rounded-3xl !px-9'
                >
                  {t('获取密钥')}
                </Button>
              </Link>
              <Link to='/pricing'>
                <Button size='large' className='!rounded-3xl !px-8'>
                  浏览模型广场
                </Button>
              </Link>
            </div>
          </div>
        </Reveal>
      </section>
    </div>
  );
};

export default Home;
