package skill

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

func LoadSkill(path string) (*Skill, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	raw := string(data)

	parts := strings.SplitN(raw, "---", 3)
	if len(parts) < 3 {
		return nil, err
	}

	var s Skill
	if err := yaml.Unmarshal([]byte(parts[1]), &s); err != nil {
		return nil, err
	}

	s.Content = parts[2]
	s.BaseDir = filepath.Dir(path)

	return &s, nil
}

// LoadReferences 加载 references 目录下的所有 .md 文件
func (s *Skill) LoadReferences() map[string]string {
	refs := make(map[string]string)
	refsDir := filepath.Join(s.BaseDir, "references")

	if info, err := os.Stat(refsDir); err == nil && info.IsDir() {
		filepath.Walk(refsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			if strings.HasSuffix(path, ".md") {
				data, _ := os.ReadFile(path)
				name := filepath.Base(path)
				refs[name] = string(data)
			}
			return nil
		})
	}

	return refs
}
