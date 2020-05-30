# [stornado.github.io](https://stornado.github.io/) [![Build Status](https://travis-ci.org/stornado/stornado.github.io.svg?branch=develop)](https://travis-ci.org/stornado/stornado.github.io)



```shell
hexo new "a new title"
hexo clean; hexo generate; hexo serve
git add source/_posts/{a-new-title}.md
git push origin develop
```



## 依赖

```json
{
    "cryptiles": "^4.1.2",
    "hawk": "^7.0.10",
    "hexo": "^4.2.0",
    "hexo-cli": "^3.1.0",
    "hexo-deployer-git": "^2.0.0",
    "hexo-filter-flowchart": "^1.0.4",
    "hexo-filter-mermaid-diagrams": "^1.0.5",
    "hexo-filter-sequence": "^1.0.3",
    "hexo-generator-archive": "^1.0.0",
    "hexo-generator-category": "^1.0.0",
    "hexo-generator-index": "^1.0.0",
    "hexo-generator-searchdb": "^1.0.8",
    "hexo-generator-tag": "^1.0.0",
    "hexo-related-popular-posts": "^4.0.0",
    "hexo-renderer-ejs": "^1.0.0",
    "hexo-renderer-pandoc": "^0.3.0",
    "hexo-renderer-stylus": "^1.1.0",
    "hexo-server": "^1.0.0",
    "hexo-symbols-count-time": "^0.7.1",
    "lodash": "^4.17.12",
    "request": "^2.68.0",
    "tunnel-agent": "^0.6.0"
}
```



```shell
git clone https://github.com/theme-next/hexo-theme-next.git themes/next
git clone https://github.com/theme-next/theme-next-pjax.git themes/next/source/lib/pjax
git clone https://github.com/theme-next/theme-next-pace.git themes/next/source/lib/pace
```



### themes/next/layout/_scripts/index.swig

```django
{%- if theme.mermaid %}
  {% include 'mermaid.swig' %}
{%- endif %}
```

### themes/next/layout/_scripts/mermaid.swig

```django
{% if theme.mermaid.enable %}
  <script src='https://unpkg.com/mermaid@{{ theme.mermaid.version }}/dist/mermaid.min.js'></script>
  <script>
    if (window.mermaid) {
      mermaid.initialize({'theme': '{{ theme.mermaid.theme }}'});
    }
  </script>
  <style>
    .mermaid {
      background: var(--body-bg-color);
    }
  </style>
{% endif %}

```

