# Suno 模型测试参数文档

## 1. `suno-music`

**非自定义模式（描述生成）**

```json
{
  "model": "suno-music",
  "version": "v5",
  "custom": false,
  "instrumental": false,
  "prompt": "深夜城市中的 lo-fi 钢琴，伴随轻柔雨声"
}
```

**自定义模式（指定风格 + 歌词）**

```json
{
  "model": "suno-music",
  "version": "v5",
  "custom": true,
  "instrumental": false,
  "title": "雨窗慢拍",
  "tags": "lo-fi, piano ballad, trip-hop, rainy night, chill",
  "prompt": "[Verse]\n我把灯关小一点\n让雨落进窗边\n你留下的那只杯子\n还在桌上转圈\n\n[Chorus]\n深夜的城，慢慢睡去\n我还醒着，等一场回忆"
}
```

## 2. `suno-lyrics`

```json
{
  "model": "suno-lyrics",
  "prompt": "写一首关于夏夜海边重逢的中文流行歌词",
  "lyrics_model": "classic"
}
```

## 3. `suno-aligned-lyrics`

```json
{
  "model": "suno-aligned-lyrics",
  "task_id": "task_source123",
  "audio_index": 1
}
```

## 4. `suno-bpm`

```json
{
  "model": "suno-bpm",
  "task_id": "task_source123",
  "audio_index": 1
}
```

## 5. `suno-concat`

```json
{
  "model": "suno-concat",
  "task_id": "task_extend123",
  "audio_index": 1
}
```

## 6. `suno-generate-video`

```json
{
  "model": "suno-generate-video",
  "task_id": "task_source123",
  "audio_index": 1
}
```

## 7. `suno-persona`

```json
{
  "model": "suno-persona",
  "task_id": "task_source123",
  "audio_index": 1,
  "name": "Warm Indie Vocal",
  "description": "温暖、自然的独立流行人声",
  "styles": "indie pop, warm vocal"
}
```

## 8. `suno-upload`

```json
{
  "model": "suno-upload",
  "audio_file_path": "https://example.com/source.mp3"
}
```

## 9. `suno-upsample-tags`

```json
{
  "model": "suno-upsample-tags",
  "tags": "lo-fi, piano, rainy night"
}
```

## 10. `suno-vox`

```json
{
  "model": "suno-vox",
  "task_id": "task_source123",
  "audio_index": 1,
  "vocal_start_s": 12,
  "vocal_end_s": 42
}
```

## 11. `suno-wav`

```json
{
  "model": "suno-wav",
  "task_id": "task_source123",
  "audio_index": 1
}
```

## 12. `suno-crop`

```json
{
  "model": "suno-crop",
  "task_id": "task_source123",
  "audio_index": 1,
  "start_s": 10,
  "end_s": 45
}
```

## 13. `suno-fade-in`

```json
{
  "model": "suno-fade-in",
  "task_id": "task_source123",
  "audio_index": 1,
  "duration_s": 4,
  "title": "Fade In Version"
}
```

## 14. `suno-fade-out`

```json
{
  "model": "suno-fade-out",
  "task_id": "task_source123",
  "audio_index": 1,
  "duration_s": 6,
  "title": "Fade Out Version"
}
```

## 15. `suno-remove-section`

```json
{
  "model": "suno-remove-section",
  "task_id": "task_source123",
  "audio_index": 1,
  "start_s": 30,
  "end_s": 42
}
```

## 16. `suno-sounds`

```json
{
  "model": "suno-sounds",
  "version": "v5.5",
  "prompt": "森林中由远及近的雷声和雨声",
  "type": "one-shot",
  "bpm": 90,
  "key": "Cm"
}
```

## 17. `suno-create-voice`

```json
{
  "model": "suno-create-voice",
  "audio_url": "https://example.com/voice.wav"
}
```

## 18. `suno-adjust-speed`

```json
{
  "model": "suno-adjust-speed",
  "task_id": "task_source123",
  "audio_index": 1,
  "speed": 1.25,
  "keep_pitch": true,
  "title": "Faster Version"
}
```

## 19. `suno-add-instrumental`

```json
{
  "model": "suno-add-instrumental",
  "version": "v5",
  "task_id": "task_source123",
  "audio_index": 1,
  "custom": true,
  "tags": "acoustic pop, warm guitar",
  "title": "Acoustic Arrangement"
}
```

## 20. `suno-add-stem`

```json
{
  "model": "suno-add-stem",
  "version": "v5.5",
  "task_id": "task_source123",
  "audio_index": 1,
  "custom": true,
  "prompt": "加入一条旋律性电吉他音轨",
  "tags": "melodic electric guitar"
}
```

## 21. `suno-add-vocals`

```json
{
  "model": "suno-add-vocals",
  "version": "v5",
  "task_id": "task_source123",
  "audio_index": 1,
  "custom": true,
  "prompt": "[Verse]\nWalking through the city lights",
  "tags": "female vocal, synth pop",
  "vocal_gender": "Female"
}
```

## 22. `suno-cover`

```json
{
  "model": "suno-cover",
  "version": "v5",
  "task_id": "task_source123",
  "audio_index": 1,
  "custom": false,
  "gpt_description": "改编为温暖的原声民谣版本"
}
```

## 23. `suno-extend`

```json
{
  "model": "suno-extend",
  "version": "v5",
  "task_id": "task_source123",
  "audio_index": 1,
  "continue_at": 90,
  "custom": true,
  "prompt": "[Chorus]\nWe will find our way home",
  "tags": "uplifting pop rock"
}
```

## 24. `suno-mashup`

```json
{
  "model": "suno-mashup",
  "version": "v5",
  "task_ids": [
    "task_source123",
    "task_source456"
  ],
  "audio_indexes": [
    1,
    2
  ],
  "custom": false,
  "gpt_description": "融合电子舞曲节奏和爵士钢琴"
}
```

## 25. `suno-midi`

```json
{
  "model": "suno-midi",
  "task_id": "task_source123",
  "audio_index": 1
}
```

## 26. `suno-remaster`

```json
{
  "model": "suno-remaster",
  "version": "v5",
  "task_id": "task_source123",
  "audio_index": 1,
  "variation_category": "normal"
}
```

## 27. `suno-replace-section`

```json
{
  "model": "suno-replace-section",
  "version": "v5",
  "task_id": "task_source123",
  "audio_index": 1,
  "start_s": 45,
  "end_s": 68,
  "infill_lyrics": "[Chorus]\nSing into the morning light",
  "tags": "anthemic pop"
}
```

## 28. `suno-sample`

```json
{
  "model": "suno-sample",
  "version": "v5",
  "task_id": "task_source123",
  "audio_index": 1,
  "start_s": 5,
  "end_s": 25,
  "instrumental": true,
  "tags": "ambient electronic"
}
```

## 29. `suno-inspo`

```json
{
  "model": "suno-inspo",
  "version": "v5",
  "audio_urls": [
    "https://example.com/reference-1.mp3",
    "https://example.com/reference-2.wav"
  ],
  "tags": "upbeat retro synth pop"
}
```

## 30. `suno-stems`

```json
{
  "model": "suno-stems",
  "task_id": "task_source123",
  "audio_index": 1,
  "stem_type": "lead_vocal"
}
```

## 31. `suno-stems-all`

```json
{
  "model": "suno-stems-all",
  "task_id": "task_source123",
  "audio_index": 1
}
```
