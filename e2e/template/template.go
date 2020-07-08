package template

const TestMasterModBusGROUP = `>>> GROUP      {{.Name}}`
const TestMasterModBusSKIP = `--- SKIP       {{.Name}} ({{.GotTime}})
{{.Skip}}`
const TestMasterModBusRUN = `=== RUN        {{.Name}}`
const TestMasterModBusPASS = `--- PASS:      {{.Name}} ({{.GotTime}})`
const TestMasterModBusFAIL = `--- FAIL:      {{.Name}} ({{.GotTime}})
{{range .Expected}}{{with .Pass}}{{else}}    {{.Name}}:

            expected: ({{.Type}}) {{.Expected}}
                 got: ({{.Type}}) {{.Got}}
{{end}}{{end}}
`
const TestSlaveModBusSKIP = `--- SKIP       {{.Name}}
{{.Skip}}`
const TestSlaveModBusRUN = `=== RUN        {{.Name}}`
const TestSlaveModBusPASS = `--- PASS:      {{.Name}}`
const TestSlaveModBusFAIL = `--- FAIL:      {{.Name}}
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

const TestSlaveCustomFATAL = `--- FATAL:      {{.Name}}`
const TestSlaveCustomSKIP = `--- SKIP       {{.Name}}
{{.Skip}}`
const TestSlaveCustomRUN = `=== RUN        {{.Name}}`
const TestSlaveCustomPASS = `--- PASS:      {{.Name}}`
const TestSlaveCustomFAIL = `--- FAIL:      {{.Name}}
{{range .Expected}}{{with .Pass}}{{else}}    {{.Name}}:

            expected: ({{.Type}}) {{.Expected}}
                 got: ({{.Type}}) {{.Got}}
{{end}}{{end}}
`
