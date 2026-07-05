// Copyright 2026 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package localization provides human-readable localized labels and guidance.
package localization

import (
	"fmt"

	"github.com/ossf/scorecard/v5/checks"
)

const (
	// LocaleDefault keeps the original English-only display.
	LocaleDefault = "en"
	// LocaleZhCN enables Chinese labels and remediation guidance.
	LocaleZhCN = "zh-CN"
)

var zhCNCheckNames = map[string]string{
	checks.CheckBinaryArtifacts:      "二进制产物",
	checks.CheckBranchProtection:     "分支保护",
	checks.CheckCIIBestPractices:     "CII 最佳实践",
	checks.CheckCITests:              "CI 测试",
	checks.CheckCodeReview:           "代码评审",
	checks.CheckContributors:         "贡献者活跃度",
	checks.CheckDangerousWorkflow:    "危险工作流",
	checks.CheckDependencyUpdateTool: "依赖更新工具",
	checks.CheckFuzzing:              "模糊测试",
	checks.CheckLicense:              "开源许可证",
	checks.CheckMaintained:           "维护状态",
	checks.CheckPackaging:            "软件包发布",
	checks.CheckPinnedDependencies:   "依赖固定",
	checks.CheckSAST:                 "静态应用安全测试",
	checks.CheckSBOM:                 "软件物料清单",
	checks.CheckSecurityPolicy:       "安全策略",
	checks.CheckSignedReleases:       "发布签名",
	checks.CheckTokenPermissions:     "Token 最小权限",
	checks.CheckVulnerabilities:      "已知漏洞",
	checks.CheckWebHooks:             "Webhook 安全",
}

var zhCNRemediations = map[string]string{
	checks.CheckBinaryArtifacts:      "避免将可执行文件、构建产物或压缩包直接提交到仓库，改用可信发布流程托管产物。",
	checks.CheckBranchProtection:     "开启默认分支保护，禁止强推和删除，要求 PR 评审后再合并。",
	checks.CheckCIIBestPractices:     "补充 OpenSSF Best Practices 信息，按项目成熟度逐步完善安全实践。",
	checks.CheckCITests:              "为主分支和 PR 配置自动化测试，确保代码合并前经过 CI 验证。",
	checks.CheckCodeReview:           "要求关键代码通过 Pull Request 或 Merge Request 评审后再合并。",
	checks.CheckContributors:         "保持维护者和贡献者活跃，避免项目长期无人响应安全问题。",
	checks.CheckDangerousWorkflow:    "审查 CI 工作流中的高危触发器和脚本注入点，避免未信任 PR 获得高权限执行。",
	checks.CheckDependencyUpdateTool: "配置 Dependabot、Renovate 或企业内部依赖更新工具，及时跟进安全补丁。",
	checks.CheckFuzzing:              "为解析器、输入处理和关键库函数补充模糊测试，降低未知崩溃风险。",
	checks.CheckLicense:              "在仓库根目录提供明确的开源许可证文件，方便企业合规评估。",
	checks.CheckMaintained:           "定期提交维护更新，及时响应 issue、PR 和安全漏洞报告。",
	checks.CheckPackaging:            "使用自动化、可审计的发布流程生成包和镜像，减少人工发布风险。",
	checks.CheckPinnedDependencies:   "固定 GitHub Actions、容器镜像和构建依赖版本，优先使用 commit hash 或 digest。",
	checks.CheckSAST:                 "在 CI 中启用静态应用安全测试，覆盖主分支和合并请求。",
	checks.CheckSBOM:                 "生成并发布 SBOM，帮助安全团队识别依赖和供应链风险。",
	checks.CheckSecurityPolicy:       "添加 SECURITY.md，说明漏洞上报方式、响应范围和处理流程。",
	checks.CheckSignedReleases:       "对 release、制品或校验文件进行签名，便于用户验证下载内容未被篡改。",
	checks.CheckTokenPermissions:     "按最小权限配置 CI Token，避免默认写权限或 write-all 权限。",
	checks.CheckVulnerabilities:      "定期扫描依赖漏洞，及时升级受影响组件并记录修复情况。",
	checks.CheckWebHooks:             "为 Webhook 配置密钥并限制接收端权限，避免伪造请求触发敏感操作。",
}

// CheckName returns a localized display name for a Scorecard check.
func CheckName(name, locale string) string {
	if locale != LocaleZhCN {
		return name
	}
	zhName, ok := zhCNCheckNames[name]
	if !ok {
		return name
	}
	return fmt.Sprintf("%s（%s）", name, zhName)
}

// Remediation returns localized remediation guidance for a Scorecard check.
func Remediation(name, locale string) string {
	if locale != LocaleZhCN {
		return ""
	}
	remediation, ok := zhCNRemediations[name]
	if !ok {
		return "参考官方文档评估该检查项，并结合国内研发流程制定整改措施。"
	}
	return remediation
}
