# Angular

## Databinding

Communication between the business logic (Typescript code) and the template (HTML)

- Output data (from Typescript to HTML)
  - String interpolation - `{{ data }}`
  - Propery binding - `[propery]="data"`
- React to User Events
  - Event binding - `(event)="expression"`
- Combination of both: **Two-way databinding**
  - `[(ngModel)]="data"`
