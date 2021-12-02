

// Code generated by "mdtogo"; DO NOT EDIT.
package generated

var SetEnforcementActionShort = `Applies the supplied enforcement action on [policy constraints](https://cloud.google.com/anthos-config-management/docs/concepts/policy-controller#constraints) within a package.`
var SetEnforcementActionLong = `
## Usage

The function will execute as follows:

1. Searches for resources with ` + "`" + `apiVersion: constraints.gatekeeper.sh/v1beta1` + "`" + `
2. Applies the enforement action value provided in KptFile to following element:
   ` + "`" + `spec.enforcementAction` + "`" + `

` + "`" + `set-enforcement-action` + "`" + ` function can be executed imperatively as follows:

  $ kpt fn eval -i gcr.io/kpt-fn/set-enforcement-action:unstable -- enforcementAction=deny

To execute ` + "`" + `set-enforcement-action` + "`" + ` declaratively include the function in kpt package pipeline as follows:
  ...
  pipeline:
    mutators:
      - image: gcr.io/kpt-fn/set-enforcement-action:unstable
        configMap:
            enforcementAction: deny
  ...
`