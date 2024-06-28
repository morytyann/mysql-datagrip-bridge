package main

import (
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const dataSourcesTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="DataSourceManagerImpl" format="xml" multifile-model="true">
    <data-source source="LOCAL" name="${name}" uuid="${uuid}" single-connection="true">
      <synchronize>true</synchronize>
      <configured-by-url>true</configured-by-url>
      <jdbc-driver>com.mysql.cj.jdbc.Driver</jdbc-driver>
      <jdbc-url>jdbc:mysql://${host}:${port}/?user=${username}&amp;password=${password}</jdbc-url>
      <time-zone>Asia/Shanghai</time-zone>
      <working-dir>$ProjectFileDir$</working-dir>
    </data-source>
  </component>
</project>
`

const dataSourcesLocalTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="dataSourceStorageLocal">
    <data-source name="${name}" uuid="${uuid}">
      <schema-mapping>
        <introspection-scope>
          <node kind="schema" negative="1" />
        </introspection-scope>
      </schema-mapping>
    </data-source>
  </component>
</project>
`

const uuid = "3ebcc32a-e11a-11ee-a857-e0e1a939a1b9"

func main() {
	var username string
	var password string
	var host string
	var port string
	// 解析命令参数
	args := os.Args
	for _, arg := range args[1:] {
		if strings.Index(arg, "-h") == 0 {
			host, _ = strings.CutPrefix(arg, "-h")
		} else if strings.Index(arg, "-P") == 0 {
			port, _ = strings.CutPrefix(arg, "-P")
		} else if strings.Index(arg, "-u") == 0 {
			username, _ = strings.CutPrefix(arg, "-u")
		} else if strings.Index(arg, "-p") == 0 {
			password, _ = strings.CutPrefix(arg, "-p")
		}
	}
	var name = `db@` + host
	// 替换参数
	var dataSourcesXml = dataSourcesTemplate
	dataSourcesXml = strings.Replace(dataSourcesXml, "${name}", name, -1)
	dataSourcesXml = strings.Replace(dataSourcesXml, "${uuid}", uuid, -1)
	dataSourcesXml = strings.Replace(dataSourcesXml, "${host}", host, -1)
	dataSourcesXml = strings.Replace(dataSourcesXml, "${port}", port, -1)
	dataSourcesXml = strings.Replace(dataSourcesXml, "${username}", username, -1)
	dataSourcesXml = strings.Replace(dataSourcesXml, "${password}", url.QueryEscape(password), -1)
	var dataSourcesLocalXml = dataSourcesLocalTemplate
	dataSourcesLocalXml = strings.Replace(dataSourcesLocalXml, "${name}", name, -1)
	dataSourcesLocalXml = strings.Replace(dataSourcesLocalXml, "${uuid}", uuid, -1)
	// 声明路径
	projectPath := filepath.Join(os.TempDir(), "mysql-datagrip-bridge")
	projectIdeaPath := filepath.Join(projectPath, ".idea")
	dataSourcesPath := filepath.Join(projectIdeaPath, "dataSources.xml")
	dataSourcesLocalPath := filepath.Join(projectIdeaPath, "dataSources.local.xml")
	if _, err := os.Stat(projectIdeaPath); os.IsNotExist(err) {
		_ = os.MkdirAll(projectIdeaPath, os.ModePerm)
	}
	// 保存配置文件
	dataSourceFile, _ := os.Create(dataSourcesPath)
	_, _ = dataSourceFile.WriteString(dataSourcesXml)
	_ = dataSourceFile.Close()
	dataSourceLocalFile, _ := os.Create(dataSourcesLocalPath)
	_, _ = dataSourceLocalFile.WriteString(dataSourcesLocalXml)
	_ = dataSourceLocalFile.Close()
	// 打开DataGrip
	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	_ = exec.Command(filepath.Join(cwd, "datagrip64.exe"), projectPath).Start()
}
