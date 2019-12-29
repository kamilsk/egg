package brew

type templateData struct {
	Name             string
	Desc             string
	Homepage         string
	Version          string
	Caveats          []string
	Plist            string
	DownloadStrategy string
	Install          []string
	Dependencies     []string
	Conflicts        []string
	Tests            []string
	CustomRequire    string
	CustomBlock      []string
	MacOS            downloadable
	Linux            downloadable
	Arm              downloadable
	Arm64            downloadable
}

type downloadable struct {
	DownloadURL string
	SHA256      string
}

const formulaTemplate = `# This file was generated by GoReleaser. DO NOT EDIT.
{{ if .CustomRequire -}}
require_relative "{{ .CustomRequire }}"
{{ end -}}
class {{ .Name }} < Formula
  desc "{{ .Desc }}"
  homepage "{{ .Homepage }}"
  version "{{ .Version }}"
  bottle :unneeded

  if OS.mac?
    {{- if .MacOS.DownloadURL }}
    url "{{ .MacOS.DownloadURL }}"
    {{- if .DownloadStrategy }}, :using => {{ .DownloadStrategy }}{{- end }}
    sha256 "{{ .MacOS.SHA256 }}"
    {{- end }}
  elsif OS.linux?
    {{- if .Linux.DownloadURL }}
    if Hardware::CPU.intel?
      url "{{ .Linux.DownloadURL }}"
      {{- if .DownloadStrategy }}, :using => {{ .DownloadStrategy }}{{- end }}
      sha256 "{{ .Linux.SHA256 }}"
    end
    {{- end }}
	{{- if or .Arm.DownloadURL .Arm64.DownloadURL }}
    if Hardware::CPU.arm?
      if Hardware::CPU.is_64_bit?
        {{- if .Arm64.DownloadURL }}
        url "{{ .Arm64.DownloadURL }}"
        {{- if .DownloadStrategy }}, :using => {{ .DownloadStrategy }}{{- end }}
        sha256 "{{ .Arm64.SHA256 }}"
        {{- end }}
      else
        {{- if .Arm.DownloadURL }}
        url "{{ .Arm.DownloadURL }}"
        {{- if .DownloadStrategy }}, :using => {{ .DownloadStrategy }}{{- end }}
        sha256 "{{ .Arm.SHA256 }}"
        {{- end }}
      end
    end
	{{- end }}
  end

  {{- with .CustomBlock }}
  {{ range $index, $element := . }}
  {{ . }}
  {{- end }}
  {{- end }}

  {{- with .Dependencies }}
  {{ range $index, $element := . }}
  depends_on "{{ . }}"
  {{- end }}
  {{- end -}}

  {{- with .Conflicts }}
  {{ range $index, $element := . }}
  conflicts_with "{{ . }}"
  {{- end }}
  {{- end }}

  def install
    {{- range $index, $element := .Install }}
    {{ . -}}
    {{- end }}
  end

  {{- with .Caveats }}

  def caveats; <<~EOS
    {{- range $index, $element := . }}
    {{ . -}}
    {{- end }}
  EOS
  end
  {{- end -}}

  {{- with .Plist }}

  plist_options :startup => false

  def plist; <<~EOS
    {{ . }}
  EOS
  end
  {{- end -}}

  {{- if .Tests }}

  test do
    {{- range $index, $element := .Tests }}
    {{ . -}}
    {{- end }}
  end
  {{- end }}
end
`
