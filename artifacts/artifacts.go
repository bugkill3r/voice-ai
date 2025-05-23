package artifacts

// import (
// 	"embed"
// 	"fmt"
// 	"io/fs"

// 	lexatic_backend "github.com/lexatic/web-backend/protos/lexatic-backend"
// 	"gopkg.in/yaml.v2"
// )

// //go:embed *.yml
// var templatesFS embed.FS

// // GetToolProviders returns a slice of ToolProvider parsed from the embedded YAML file
// func GetToolProviders() ([]*lexatic_backend.ToolProvider, error) {
// 	// Read all .yml files in the embedded filesystem
// 	files, err := fs.Glob(templatesFS, "*.yml")
// 	if err != nil {
// 		return nil, fmt.Errorf("error globbing YAML files: %w", err)
// 	}
// 	var allToolProviders []*lexatic_backend.ToolProvider

// 	for _, file := range files {
// 		// Read the content of each YAML file
// 		content, err := templatesFS.ReadFile(file)
// 		if err != nil {
// 			return nil, fmt.Errorf("error reading file %s: %w", file, err)
// 		}

// 		// Parse the YAML content
// 		var yamlToolProviders []struct {
// 			ID                   uint64                   `yaml:"id"`
// 			ToolID               uint64                   `yaml:"tool_id"`
// 			Provider             lexatic_backend.Provider `yaml:"provider"`
// 			Name                 string                   `yaml:"name"`
// 			Feature              []string                 `yaml:"feature"`
// 			ConnectConfiguration map[string]string        `yaml:"connect_configuration"`
// 			Actor                string                   `yaml:"actor"`
// 		}

// 		err = yaml.Unmarshal(content, &yamlToolProviders)
// 		if err != nil {
// 			return nil, fmt.Errorf("error unmarshaling YAML from %s: %w", file, err)
// 		}
// 		for _, ytp := range yamlToolProviders {
// 			tp := &lexatic_backend.ToolProvider{
// 				Id:                   ytp.ID,
// 				Provider:             &ytp.Provider,
// 				Name:                 ytp.Name,
// 				Feature:              ytp.Feature,
// 				Actor:                ytp.Actor,
// 				ConnectConfiguration: ytp.ConnectConfiguration,
// 			}
// 			allToolProviders = append(allToolProviders, tp)
// 		}
// 	}

// 	return allToolProviders, nil
// }
