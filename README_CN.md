# OpenSSF Scorecard 国内化二次开发版

本项目基于 [OpenSSF Scorecard](https://github.com/ossf/scorecard) 进行二次开发，目标是面向国内企业、开源社区和私有化研发环境，提供更容易落地的开源供应链安全检查工具。

原项目用于自动评估开源项目的安全实践，例如分支保护、CI 测试、危险工作流、依赖固定、漏洞管理、安全策略、发布签名等。本项目将在保留原有能力的基础上，逐步增加中文报告、国内代码托管平台适配、国内 CI/CD 示例和本土化规则说明。

## 项目定位

- 面向国内开发团队的开源供应链安全评分工具
- 基于 OpenSSF Scorecard 的合规二次开发项目
- 支持学习、企业内部审计、开源项目健康度评估和简历项目展示

## 第一阶段目标

1. 提供中文 README 和使用说明。
2. 梳理 Scorecard 检查项在国内企业研发流程中的含义。
3. 增加 Gitee、私有 GitLab、Jenkins、Gitee Go 等场景的落地文档。
4. 设计中文报告输出能力，方便安全团队、研发负责人和项目维护者阅读。
5. 保留原项目许可证、版权声明和来源说明，避免不合规复制。

## 推荐使用场景

### 检查 GitHub 项目

```powershell
go run main.go --repo=github.com/ossf/scorecard --show-details
```

### 输出 JSON 报告

```powershell
go run main.go --repo=github.com/ossf/scorecard --format=json
```

### 检查指定规则

```powershell
go run main.go --repo=github.com/ossf/scorecard --checks=Pinned-Dependencies,Security-Policy
```

### 输出中文检查项名称和整改建议

```powershell
go run main.go --repo=github.com/ossf/scorecard --show-details --locale=zh-CN
```

启用后，默认文本报告中的检查项名称会显示为 `Branch-Protection（分支保护）` 这类中英文组合，并新增 `中文整改建议` 列，例如：

```text
Branch-Protection（分支保护）
中文整改建议：开启默认分支保护，禁止强推和删除，要求 PR 评审后再合并。
```

JSON 输出仍保留原始字段，避免破坏自动化集成。

## 国内化改造方向

### 中文化报告

将原始英文检查项解释为中文，例如：

| 原检查项 | 中文说明 | 国内企业关注点 |
| --- | --- | --- |
| Branch-Protection | 分支保护 | 防止主分支被误删、强推或绕过评审 |
| Dangerous-Workflow | 危险 CI 工作流 | 防止 PR 触发高权限脚本造成供应链攻击 |
| Pinned-Dependencies | 依赖固定 | 防止构建过程拉取到被篡改的 action、镜像或包 |
| Security-Policy | 安全漏洞披露策略 | 建立漏洞上报和响应机制 |
| Token-Permissions | CI Token 最小权限 | 降低 CI 凭证被滥用风险 |

### 国内平台适配

后续可逐步支持或增强：

- Gitee 仓库基础信息检查
- 私有 GitLab 仓库检查
- Jenkins Pipeline 安全建议
- Gitee Go / GitLab CI / GitHub Actions 示例模板
- 面向内网环境的离线扫描和报告导出

## 合规说明

本项目是 OpenSSF Scorecard 的衍生项目。原项目采用 Apache License 2.0，本项目保留原 LICENSE、版权声明和源码中的许可证头。详细说明见 [DERIVATIVE_WORK.md](DERIVATIVE_WORK.md)。

## 简历描述示例

> 基于 OpenSSF Scorecard 二次开发开源供应链安全评分工具，面向国内企业研发环境补充中文报告、Gitee/私有 GitLab 适配规划、CI/CD 安全检查说明和本土化落地文档，覆盖分支保护、依赖固定、危险工作流、Token 最小权限和安全策略等供应链风险。
