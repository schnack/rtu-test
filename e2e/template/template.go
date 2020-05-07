package template

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

const TestSlaveSkip = `--- SKIP       {{.Name}}
{{.Skip}}`
const TestSlaveRUN = `=== RUN        {{.Name}}`
const TestSlavePASS = `--- PASS:      {{.Name}}`
const TestSlaveFAIL = `--- FAIL:      {{.Name}}
{{range .ExpectedCoils}}{{with .Pass}}{{else}}    {{.Name}}:

            expected: ({{.Type}}) {{.Expected}}
                 got: ({{.Type}}) {{.Got}}
{{end}}{{end}}
{{range .ExpectedDiscreteInput}}{{with .Pass}}{{else}}    {{.Name}}:

            expected: ({{.Type}}) {{.Expected}}
                 got: ({{.Type}}) {{.Got}}
{{end}}{{end}}
{{range .ExpectedHoldingRegisters}}{{with .Pass}}{{else}}    {{.Name}}:

            expected: ({{.Type}}) {{.Expected}}
                 got: ({{.Type}}) {{.Got}}
{{end}}{{end}}
{{range .ExpectedInputRegisters}}{{with .Pass}}{{else}}    {{.Name}}:

            expected: ({{.Type}}) {{.Expected}}
                 got: ({{.Type}}) {{.Got}}
{{end}}{{end}}
`
