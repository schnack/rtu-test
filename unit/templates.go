package unit

const TestGROUP = `>>> GROUP      {{.Name}}`
const TestSKIP = `--- SKIP       {{.Name}} ({{.GotTime}})
{{.Skip}}`
const TestRUN = `=== RUN        {{.Name}}`
const TestPASS = `--- PASS:      {{.Name}} ({{.GotTime}})`
const TestFAIL = `--- FAIL:      {{.Name}} ({{.GotTime}})
{{range .Expected}}{{with .Pass}}{{else}}    {{.Name}}:

            expected: ({{.Type}}) {{.Expected}}
                 got: ({{.Type}}) {{.Got}}
{{end}}{{end}}
`
