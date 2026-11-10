package common

import "testing"

func TestNormalizeImageFields(t *testing.T) {
	cases := []struct {
		name    string
		model   string
		in      ImageFields
		strict  bool
		wantErr bool
		want    ImageFields
	}{
		{
			name:  "gpt size only",
			model: "gpt-image-2",
			in:    ImageFields{Size: "1536X1024", Quality: "High"},
			want:  ImageFields{Size: "1536x1024", Quality: "high"},
		},
		{
			name:    "gpt size with ratio",
			model:   "gpt-image-2",
			in:      ImageFields{Size: "1024x1024", Ratio: "1:1"},
			wantErr: true,
		},
		{
			name:    "gpt legacy quality",
			model:   "OG-image-2",
			in:      ImageFields{Quality: "hd"},
			wantErr: true,
		},
		{
			name:  "gpt xhigh alone",
			model: "gpt-image-2",
			in:    ImageFields{Quality: "xhigh"},
			want:  ImageFields{Quality: "xhigh"},
		},
		{
			name:    "gpt ratio as size",
			model:   "gpt-image-2",
			in:      ImageFields{Size: "16:9"},
			wantErr: true,
		},
		{
			name:  "gemini resolution and ratio",
			model: "gemini-2.5-flash-image",
			in:    ImageFields{Resolution: "2k", Ratio: "16:9"},
			want:  ImageFields{Resolution: "2K", Ratio: "16:9"},
		},
		{
			name:    "gemini rejects size",
			model:   "gemini-3-pro-image",
			in:      ImageFields{Size: "1024x1024"},
			wantErr: true,
		},
		{
			name:   "public gg pro accepts resolution and ratio",
			model:  "GG-3-pro-image",
			strict: true,
			in:     ImageFields{Resolution: "1K", Ratio: "1:1"},
			want:   ImageFields{Resolution: "1K", Ratio: "1:1"},
		},
		{
			name:   "public gg 3.1 accepts 720p",
			model:  "GG-3.1-flash-image",
			strict: true,
			in:     ImageFields{Resolution: "720p", Ratio: "1:8"},
			want:   ImageFields{Resolution: "720P", Ratio: "1:8"},
		},
		{
			name:    "gemini 2.5 rejects 720p",
			model:   "gemini-2.5-flash-image",
			in:      ImageFields{Resolution: "720P"},
			wantErr: true,
		},
		{
			name:  "gemini 3.1 720p and extreme ratio",
			model: "gemini-3.1-flash-image",
			in:    ImageFields{Resolution: "720p", Ratio: "1:8"},
			want:  ImageFields{Resolution: "720P", Ratio: "1:8"},
		},
		{
			name:    "kling 2.1 rejects 4k",
			model:   "kling-image-2.1",
			in:      ImageFields{Resolution: "4k"},
			wantErr: true,
		},
		{
			name:  "kling omni auto and 4k",
			model: "kling-image-o1",
			in:    ImageFields{Resolution: "4K", Ratio: "auto"},
			want:  ImageFields{Resolution: "4k", Ratio: "auto"},
		},
		{
			name:    "kling scene rejects ratio",
			model:   "kling-image-scene",
			in:      ImageFields{Ratio: "1:1"},
			wantErr: true,
		},
		{
			name:  "vidu canonical resolution",
			model: "vidu-image-q2",
			in:    ImageFields{Resolution: "1080P", Ratio: "9:16"},
			want:  ImageFields{Resolution: "1080p", Ratio: "9:16"},
		},
		{
			name:    "qwen pixels too small",
			model:   "qwen-image-0925",
			in:      ImageFields{Size: "256x256"},
			wantErr: true,
		},
		{
			name:  "qwen pixels ok",
			model: "qwen-image-0925",
			in:    ImageFields{Size: "1024x1024"},
			want:  ImageFields{Size: "1024x1024"},
		},
		{
			name:    "hunyuan product too large",
			model:   "hunyuan-image-3.0",
			in:      ImageFields{Size: "2048x1024"},
			wantErr: true,
		},
		{
			name:    "seedream rejects unknown dated pro id",
			model:   "doubao-seedream-5-0-pro-999999",
			in:      ImageFields{Resolution: "2K"},
			strict:  true,
			wantErr: true,
		},
		{
			name:    "seedream rejects unlisted ratio",
			model:   "doubao-seedream-5.0-pro",
			in:      ImageFields{Resolution: "2K", Ratio: "5:4"},
			wantErr: true,
		},
		{
			name:  "seedream public pro",
			model: "doubao-seedream-5.0-pro",
			in:    ImageFields{Resolution: "1.5K", Ratio: "16:9"},
			want:  ImageFields{Resolution: "1.5K", Ratio: "16:9"},
		},
		{
			name:    "seedream public flash rejects 4k",
			model:   "doubao-seedream-5.0-flash",
			in:      ImageFields{Resolution: "4K"},
			wantErr: true,
		},
		{
			name:  "seedream public lite",
			model: "doubao-seedream-5.0-lite",
			in:    ImageFields{Resolution: "3K"},
			want:  ImageFields{Resolution: "3K"},
		},
		{
			name:  "seedream pro 1.5k with ratio",
			model: "doubao-seedream-5-0-pro-260628",
			in:    ImageFields{Resolution: "1.5k", Ratio: "16:9"},
			want:  ImageFields{Resolution: "1.5K", Ratio: "16:9"},
		},
		{
			name:    "seedream pro rejects 4k",
			model:   "doubao-seedream-5-0-pro-260628",
			in:      ImageFields{Resolution: "4K"},
			wantErr: true,
		},
		{
			name:  "seedream lite 4k",
			model: "doubao-seedream-5-0-260128",
			in:    ImageFields{Resolution: "4K"},
			want:  ImageFields{Resolution: "4K"},
		},
		{
			name:    "seedream lite rejects 1k",
			model:   "doubao-seedream-5-0-t2i-250804",
			in:      ImageFields{Resolution: "1K"},
			wantErr: true,
		},
		{
			name:    "seedream size with ratio",
			model:   "doubao-seedream-4-5-251128",
			in:      ImageFields{Size: "2048x2048", Ratio: "1:1"},
			wantErr: true,
		},
		{
			name:  "seedream 4.0 1k",
			model: "seedream-4-0-250828",
			in:    ImageFields{Resolution: "1K"},
			want:  ImageFields{Resolution: "1K"},
		},
		{
			name:  "seedream 3.0 size only",
			model: "doubao-seedream-3-0-t2i-250415",
			in:    ImageFields{Size: "1024x1024"},
			want:  ImageFields{Size: "1024x1024"},
		},
		{
			name:    "seedream 3.0 rejects resolution",
			model:   "doubao-seedream-3-0-t2i-250415",
			in:      ImageFields{Resolution: "2K"},
			wantErr: true,
		},
		{
			name:    "apimart n too large",
			model:   "gpt-image-2.6-flare",
			in:      ImageFields{N: 5, Quality: "max"},
			wantErr: true,
		},
		{
			name:    "apimart gif",
			model:   "gpt-image-2.5-sunburst",
			in:      ImageFields{OutputFormat: "gif"},
			wantErr: true,
		},
		{
			name:    "apimart size with resolution",
			model:   "gpt-image-2.6-sunburst",
			in:      ImageFields{Size: "1024x1024", Resolution: "2K"},
			wantErr: true,
		},
		{
			name:  "flux resolution alias",
			model: "flux-3-image",
			in:    ImageFields{Resolution: "768", Ratio: "AUTO"},
			want:  ImageFields{Resolution: "768sq", Ratio: "auto"},
		},
		{
			name:  "flux 1.5k",
			model: "flux-3-image-klein",
			in:    ImageFields{Resolution: "1.5K", Ratio: "7:5", OutputFormat: "JPG"},
			want:  ImageFields{Resolution: "1.5k", Ratio: "7:5", OutputFormat: "jpeg"},
		},
		{
			name:    "flux rejects size",
			model:   "flux-3-image",
			in:      ImageFields{Size: "1024x1024"},
			wantErr: true,
		},
		{
			name:    "flux rejects quality",
			model:   "flux-3-image-klein",
			in:      ImageFields{Quality: "high"},
			wantErr: true,
		},
		{
			name:    "flux n too large",
			model:   "flux-3-image",
			in:      ImageFields{N: 2},
			wantErr: true,
		},
		{
			name:    "flux klein rejects webp",
			model:   "flux-3-image-klein",
			in:      ImageFields{OutputFormat: "webp"},
			wantErr: true,
		},
		{
			name:    "flux rejects unknown ratio",
			model:   "flux-3-image",
			in:      ImageFields{Ratio: "3:1"},
			wantErr: true,
		},
		{
			name:  "mai ratio and resolution",
			model: "mai-image-2.6",
			in:    ImageFields{Resolution: "2k", Ratio: "16:9"},
			want:  ImageFields{Resolution: "2K", Ratio: "16:9"},
		},
		{
			name:    "mai rejects 4k",
			model:   "mai-image-2.6-flash",
			in:      ImageFields{Resolution: "4K"},
			wantErr: true,
		},
		{
			name:    "mai pixel too small",
			model:   "mai-image-2.6",
			in:      ImageFields{Size: "512x512"},
			wantErr: true,
		},
		{
			name:  "qwen image 3 size",
			model: "qwen-image-3.0-pro",
			in:    ImageFields{Size: "1600x900", N: 6},
			want:  ImageFields{Size: "1600x900", N: 6},
		},
		{
			name:    "qwen image 3 n too large",
			model:   "qwen-image-3.0",
			in:      ImageFields{N: 7},
			wantErr: true,
		},
		{
			name:    "muse rejects resolution",
			model:   "muse-image-ext",
			in:      ImageFields{Resolution: "1K"},
			wantErr: true,
		},
		{
			name:  "muse ratio",
			model: "muse-image-ext",
			in:    ImageFields{Ratio: "3:2"},
			want:  ImageFields{Ratio: "3:2"},
		},
		{
			name:  "hy resolution",
			model: "hy-image-v3.5-preview",
			in:    ImageFields{Resolution: "4k", Ratio: "auto"},
			want:  ImageFields{Resolution: "4K", Ratio: "auto"},
		},
		{
			name:    "hy rejects n",
			model:   "hy-image-v3.5-preview",
			in:      ImageFields{N: 2},
			wantErr: true,
		},
		{
			name:   "v1 unknown model skipped",
			model:  "dall-e-3",
			in:     ImageFields{Size: "1024x1024", Quality: "standard"},
			strict: false,
			want:   ImageFields{Size: "1024x1024", Quality: "standard"},
		},
		{
			name:    "v2 unknown model rejected",
			model:   "dall-e-3",
			in:      ImageFields{Size: "1024x1024"},
			strict:  true,
			wantErr: true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeImageFields(tt.model, tt.in, tt.strict)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestImageUpstreamDefaultQuality(t *testing.T) {
	if got := ImageUpstreamDefaultQuality("gpt-image-2.5-flare"); got != "low" {
		t.Fatalf("apimart default quality = %q", got)
	}
	if got := ImageUpstreamDefaultQuality("gpt-image-2"); got != "" {
		t.Fatalf("gpt-image-2 default quality = %q", got)
	}
	if got := ImageRefLimit("gpt-image-2.6-flare"); got != 16 {
		t.Fatalf("ref limit = %d", got)
	}
	if got := ImageUpstreamDefaultQuality("flux-3-image"); got != "" {
		t.Fatalf("flux default quality = %q", got)
	}
	if got := ImageRefLimit("flux-3-image-klein"); got != 10 {
		t.Fatalf("flux ref limit = %d", got)
	}
}
