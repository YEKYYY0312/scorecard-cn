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

package localization

import "testing"

func TestLocalizeCheckNameZhCN(t *testing.T) {
	t.Parallel()

	got := CheckName("Branch-Protection", LocaleZhCN)
	want := "Branch-Protection（分支保护）"
	if got != want {
		t.Fatalf("CheckName() = %q, want %q", got, want)
	}
}

func TestCheckNameKeepsDefaultLocaleUnchanged(t *testing.T) {
	t.Parallel()

	got := CheckName("Branch-Protection", LocaleDefault)
	want := "Branch-Protection"
	if got != want {
		t.Fatalf("CheckName() = %q, want %q", got, want)
	}
}

func TestRemediationZhCN(t *testing.T) {
	t.Parallel()

	got := Remediation("Branch-Protection", LocaleZhCN)
	want := "开启默认分支保护，禁止强推和删除，要求 PR 评审后再合并。"
	if got != want {
		t.Fatalf("Remediation() = %q, want %q", got, want)
	}
}

func TestRemediationDefaultLocaleIsEmpty(t *testing.T) {
	t.Parallel()

	if got := Remediation("Branch-Protection", LocaleDefault); got != "" {
		t.Fatalf("Remediation() = %q, want empty string", got)
	}
}

func TestUnknownCheckUsesFallbackRemediation(t *testing.T) {
	t.Parallel()

	got := Remediation("Unknown-Check", LocaleZhCN)
	want := "参考官方文档评估该检查项，并结合国内研发流程制定整改措施。"
	if got != want {
		t.Fatalf("Remediation() = %q, want %q", got, want)
	}
}
