# 衍生项目与许可证说明

本仓库基于 OpenSSF Scorecard 进行二次开发。

上游项目：

- 名称：OpenSSF Scorecard
- 地址：https://github.com/ossf/scorecard
- 许可证：Apache License 2.0

## 本项目的原则

1. 保留上游项目的 `LICENSE` 文件。
2. 保留源代码文件中的版权和许可证声明。
3. 明确说明本项目来源于 OpenSSF Scorecard。
4. 对新增文档、功能和本土化改动进行清晰记录。
5. 不将上游项目伪装为完全原创项目。

## 当前新增内容

本项目当前新增的国内化内容包括：

- `README_CN.md`：中文项目介绍、使用方式和国内化方向。
- `docs/china-localization-roadmap.md`：国内化二次开发路线图。

## 后续改动记录模板

每次新增功能时，建议按如下格式记录：

```text
日期：
改动类型：文档 / 功能 / 测试 / CI / 重构
改动内容：
影响范围：
是否修改上游核心逻辑：
```

## 简历和作品集建议

推荐表述：

> 基于 OpenSSF Scorecard 的 Apache-2.0 许可证进行合规二次开发，面向国内研发环境补充中文化报告、国内代码托管平台适配规划和 CI/CD 安全实践文档。

不推荐表述：

> 独立开发了一个和 OpenSSF Scorecard 完全一样的工具。

