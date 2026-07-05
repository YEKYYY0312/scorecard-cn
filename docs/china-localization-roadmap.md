# 国内化二次开发路线图

本文档记录基于 OpenSSF Scorecard 的国内化改造计划。目标不是简单复制上游项目，而是在合规保留上游能力的基础上，增加国内企业更容易使用、理解和落地的能力。

## 阶段 1：项目定位和中文文档

状态：进行中

目标：

- 增加中文 README。
- 说明项目来源、许可证和二次开发边界。
- 梳理适合国内企业阅读的检查项说明。
- 提供基础命令示例。

交付物：

- `README_CN.md`
- `DERIVATIVE_WORK.md`
- `docs/china-localization-roadmap.md`

## 阶段 2：中文报告输出

状态：进行中

目标：

- 为常见检查项提供中文名称和中文解释。
- 在 JSON 或文本输出中增加中文说明字段。
- 保留英文原始字段，避免破坏上游兼容性。
- 当前已支持在默认文本输出中通过 `--locale=zh-CN` 展示检查项中文名称。

优先检查项：

| 检查项 | 中文名称 | 优先级 |
| --- | --- | --- |
| Branch-Protection | 分支保护 | 高 |
| Dangerous-Workflow | 危险工作流 | 高 |
| Token-Permissions | Token 最小权限 | 高 |
| Pinned-Dependencies | 依赖固定 | 高 |
| Security-Policy | 安全策略 | 中 |
| Signed-Releases | 发布签名 | 中 |
| Vulnerabilities | 已知漏洞 | 中 |

设计原则：

- 不直接删除或改名上游字段。
- 通过新增映射层实现中文展示。
- 保持命令行默认行为兼容上游。

可能命令形式：

```powershell
go run main.go --repo=github.com/ossf/scorecard --show-details --locale=zh-CN
```

## 阶段 3：国内代码托管平台适配

状态：计划中

目标：

- 调研 Gitee API 能否支持仓库、分支、提交、PR、Issue、CI 配置等元数据读取。
- 设计 Gitee 客户端接口，不直接耦合到现有 GitHub 客户端。
- 优先支持只读检查，降低权限和安全风险。

候选能力：

- 仓库是否归档或长期未维护。
- 是否存在 LICENSE。
- 是否存在 SECURITY.md 或等效安全策略。
- 是否存在 CI 配置。
- 默认分支是否启用保护规则。

## 阶段 4：国内 CI/CD 安全模板

状态：计划中

目标：

- 提供 GitLab CI、Jenkins、Gitee Go 的安全示例。
- 说明依赖固定、最小权限、密钥管理和构建隔离。
- 将 Scorecard 检查项翻译成国内研发团队能执行的整改动作。

示例模板：

```text
ci-examples/
  gitlab-ci/
  jenkins/
  gitee-go/
```

## 阶段 5：企业落地报告

状态：计划中

目标：

- 生成适合研发负责人和安全团队阅读的 Markdown/HTML 报告。
- 按风险等级排序问题。
- 输出整改建议、影响范围和参考命令。

报告结构建议：

```text
项目概览
总体评分
高风险问题
中风险问题
整改建议
检查项详情
附录：原始 Scorecard 输出
```

## 简历项目亮点

可以在简历中突出：

- Go 开源项目二次开发经验。
- 开源供应链安全理解。
- Apache-2.0 许可证合规意识。
- GitHub/GitLab/Gitee 代码托管平台差异分析。
- CI/CD 安全实践和自动化扫描能力。
- 中文化工程落地能力。
