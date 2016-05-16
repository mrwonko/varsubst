# varsubst
Reads from stdin, writes to stdout. Substitutes all ${VAR} occurances with their set values. 
I am using it to parameterize Kubernetes configuration files. I did not want to use gettext#envsubst.

Like this:
```bash
> varsubst < rc-controller.tmpl.yaml > rc-controller.yaml
```
