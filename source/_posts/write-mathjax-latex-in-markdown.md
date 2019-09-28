---
title: Write MathJax/LaTeX in Markdown
date: 2019-09-28 12:19:30
mathjax: true
tags:
  - markdown
  - mathjax
  - latex
  - tutorial
categories:
  - tutorial
---

在markdown中编写公式如：

$$f(x)=\frac{1}{\sqrt{2 \pi \sigma x}} e^{-\frac{(x-\mu)^{2}}{2 \sigma^{2}}}$$

## 语法介绍

插入公式使用 `$LaTeX语法表达式$` 或 `$$LaTeX语法表达式$$`

```markdown
求证: $2013^2 + 2013^2 \times 2014^2 + 2014^2$ 是完全平方数
```

> 求证 $2013^2 + 2013^2 \times 2014^2 + 2014^2$ 是完全平方数

```markdown
已知:
$$\begin{cases} a + b + c = 1 \\ a^2 + b^2 + c^2 = 2 \\ a^3 + b^3 + c^3 = 3\end{cases}$$
求 $abc$ 的值
```

> 已知:
> $$\begin{cases} a + b + c = 1 \\ a^2 + b^2 + c^2 = 2 \\ a^3 + b^3 + c^3 = 3\end{cases}$$
>
> 求 $abc$ 的值

### 分数

分数使用 `\frac{分子}{分母}` 或者 `\cfrac`

```markdown
$$\frac{1}{3}$$ 以及 $$\cfrac{1}{3}$$
```

> $$\frac{1}{3}$$ 以及 $$\cfrac{1}{3}$$

### 上下标

使用 `^` 来表示上标，`_` 来表示下标，同时如果上下标的内容多于一个字符，可以使用 `{}` 来将这些内容括起来当做一个整体。
与此同时，上下标是可以嵌套的。

```markdown
$$x = a_{1}^n + a_{2}^n + a_{3}^n$$
```

> $$x = a_{1}^n + a_{2}^n + a_{3}^n$$

如果希望左右两边都能有上下标，可以使用 `\sideset` 语法

```markdown
$$\sideset{^1_2}{^3_4}X$$

$$\sideset{^{11}_{22}}{^{33}_{44}}X$$
```

> $$\sideset{^1_2}{^3_4}X$$
>
> $$\sideset{^{11}_{22}}{^{33}_{44}}X$$

### 括号

`()`，`[]` 和 `|` 都表示它们自己，但是 `{}` 因为有特殊作用因此当需要显示大括号时一般使用 `\lbrace \rbrace` 来表示。

```markdown
$$f(x, y) = 100 * \lbrace[(x + y) * 3] - 5 + |a|\rbrace$$
```

> $$f(x, y) = 100 * \lbrace[(x + y) * 3] - 5 + |a|\rbrace$$

### 开方

开方使用 `\sqrt[次数]{被开方数}` 这样的语法

```markdown
$$\sqrt[3]{X}$$
$$\sqrt{5 - x}$$
```

> $$\sqrt[3]{X}$$
> $$\sqrt{5 - x}$$

### 方程组

使用 `$$\begin{cases} LaTex语法块 \end{cases}$$`

```markdown
$$\begin{cases}{\oiint_{\mathrm{S}} E \cdot d s=\frac{Q}{\varepsilon_{0}}} \\ {\oiint_{\mathrm{S}} B \cdot d s=0} \\ {\oint_{L} E \cdot d \ell=-\frac{\mathrm{d} \Phi_{B}}{d t}} \\ {\oint_{L} B \cdot d \ell=\mu_{0} \mathrm{I}+\mu_{0} \varepsilon_{0} \frac{d \Phi_{E}}{d t}}\end{cases}$$
```

> $$\begin{cases}{\oiint_{\mathrm{S}} E \cdot d s=\frac{Q}{\varepsilon_{0}}} \\ {\oiint_{\mathrm{S}} B \cdot d s=0} \\ {\oint_{L} E \cdot d \ell=-\frac{\mathrm{d} \Phi_{B}}{d t}} \\ {\oint_{L} B \cdot d \ell=\mu_{0} \mathrm{I}+\mu_{0} \varepsilon_{0} \frac{d \Phi_{E}}{d t}}\end{cases}$$

或使用 `$$\begin{array} LaTex语法块 \end{array}$$`

```markdown
$$\begin{array}{c}{\oiint_{\mathrm{S}} E \cdot d s=\frac{Q}{\varepsilon_{0}}} \\ {\oiint_{\mathrm{S}} B \cdot d s=0} \\ {\oint_{L} E \cdot d \ell=-\frac{\mathrm{d} \Phi_{B}}{d t}} \\ {\oint_{L} B \cdot d \ell=\mu_{0} \mathrm{I}+\mu_{0} \varepsilon_{0} \frac{d \Phi_{E}}{d t}}\end{array}$$
```

> $$\begin{array}{c}{\oiint_{\mathrm{S}} E \cdot d s=\frac{Q}{\varepsilon_{0}}} \\ {\oiint_{\mathrm{S}} B \cdot d s=0} \\ {\oint_{L} E \cdot d \ell=-\frac{\mathrm{d} \Phi_{B}}{d t}} \\ {\oint_{L} B \cdot d \ell=\mu_{0} \mathrm{I}+\mu_{0} \varepsilon_{0} \frac{d \Phi_{E}}{d t}}\end{array}$$

### 换行

使用 `\\` 换行， 使用 `\,\!` 取消换行

> $$f(x) \,\!
> = \sum_{n=0}^\infty a_n x^n \\
> = a_0+a_1x+a_2x^2+\cdots$$



## 字符及运算符

### 希腊字母

|    代码    |    大写    |    代码    |    小写    |
| :--------: | :--------: | :--------: | :--------: |
|    `A`     |    $A$     |  `\alpha`  |  $\alpha$  |
|    `B`     |    $B$     |  `\beta`   |  $\beta$   |
|  `\Gamma`  |  $\Gamma$  |  `\gamma`  |  $\gamma$  |
|  `\Delta`  |  $\Delta$  |  `\delta`  |  $\delta$  |
|    `E`     |    $E$     | `\epsilon` | $\epsilon$ |
|    `Z`     |    $Z$     |  `\zeta`   |  $\zeta$   |
|    `H`     |    $H$     |   `\eta`   |   $\eta$   |
|  `\Theta`  |  $\Theta$  |  `\theta`  |  $\theta$  |
|    `I`     |    $I$     |  `\iota`   |  $\iota$   |
|    `K`     |    $K$     |  `\kappa`  |  $\kappa$  |
| `\Lambda`  | $\Lambda$  | `\lambda`  | $\lambda$  |
|    `M`     |    $M$     |   `\mu`    |   $\mu$    |
|    `N`     |    $N$     |   `\nu`    |   $\nu$    |
|   `\Xi`    |   $\Xi$    |   `\xi`    |   $\xi$    |
|    `O`     |    $O$     | `\omicron` | $\omicron$ |
|   `\Pi`    |   $\Pi$    |   `\pi`    |   $\pi$    |
|    `P`     |    $P$     |   `\rho`   |   $\rho$   |
|  `\Sigma`  |  $\Sigma$  |  `\sigma`  |  $\sigma$  |
|    `T`     |    $T$     |   `\tau`   |   $\tau$   |
| `\Upsilon` | $\Upsilon$ | `\upsilon` | $\upsilon$ |
|   `\Phi`   |   $\Phi$   |   `\phi`   |   $\phi$   |
|    `X`     |    $X$     |   `\chi`   |   $\chi$   |
|   `\Psi`   |   $\Psi$   |   `\psi`   |   $\psi$   |
|  `\Omega`  |  $\Omega$  |  `\omega`  |  $\omega$  |

### 其他字符

#### 关系运算符

|     符号     | 代码         |
| :----------: | :----------- |
|    $\pm$     | `\pm`        |
|   $\times$   | `\times`     |
|    $\div$    | `\div`       |
|    $\mid$    | `\mid`       |
|   $\nmid$    | `\nmid`      |
|   $\cdot$    | `\cdot`      |
|   $\circ$    | `\circ`      |
|    $\ast$    | `\ast`       |
|  $\bigodot$  | `\bigodot`   |
| $\bigotimes$ | `\bigotimes` |
| $\bigoplus$  | `\bigoplus`  |
|    $\leq$    | `\leq`       |
|    $\geq$    | `\geq`       |
|    $\neq$    | `\neq`       |
|  $\approx$   | `\approx`    |
|   $\equiv$   | `\equiv`     |
|    $\sum$    | `\sum`       |
|   $\prod$    | `\prod`      |
|  $\coprod$   | `\coprod`    |

#### 集合运算符

|    符号     | 代码        |
| :---------: | :---------- |
| $\emptyset$ | `\emptyset` |
|    $\in$    | `\in`       |
|  $\notin$   | `\notin`    |
|  $\subset$  | `\subset`   |
|  $\supset$  | `\supset`   |
| $\subseteq$ | `\subseteq` |
| $\supseteq$ | `\supseteq` |
|  $\bigcap$  | `\bigcap`   |
|  $\bigcup$  | `\bigcup`   |
|  $\bigvee$  | `\bigvee`   |
| $\bigwedge$ | `\bigwedge` |
| $\biguplus$ | `\biguplus` |
| $\bigsqcup$ | `\bigsqcup` |

#### 对数运算符

|  符号  | 代码   |
| :----: | :----- |
| $\log$ | `\log` |
| $\lg$  | `\lg`  |
| $\ln$  | `\ln`  |

#### 三角运算符

|   符号   | 代码     |
| :------: | :------- |
|  $\bot$  | `\bot`   |
| $\angle$ | `\angle` |
|  $\sin$  | `\sin`   |
|  $\cos$  | `\cos`   |
|  $\tan$  | `\tan`   |
|  $\cot$  | `\cot`   |
|  $\sec$  | `\sec`   |
|  $\csc$  | `\csc`   |

#### 微积分运算符

|     符号     | 代码         |
| :----------: | :----------- |
|   $\prime$   | `\prime`     |
|    $\int$    | `\int`       |
|   $\iint$    | `\iint`      |
|   $\iiint$   | `\iiint`     |
|  $\iiiint$   | `\iiiint`    |
|   $\oint$    | `\oint`      |
|   $\oiint$   | `\oiint`     |
|    $\lim$    | `\lim`       |
|   $\infty$   | `\infty`     |
|   $\nabla$   | `\nabla`     |
| $\mathrm{d}$ | `\mathrm{d}` |

## 软件推荐

1. 截图生成LaTeX [Mathpix Snip](https://mathpix.com/)
2. 小巧易用的Markdown编辑器 [Typora](https://typora.io/)
