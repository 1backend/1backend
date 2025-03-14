# page component

## Usage

## One column - only icons + content

```html
<app-page
	title="Title"
	[menuWidth]="'3rem'" <!-- Overriding default which is 90% -->
	[columnWidths]="['3rem']"
	[mobileColumnWidths]="['3rem']"
>
	<ng-template #columnContent>
		<div>
			<app-icon-menu></app-icon-menu>
		</div>
	</ng-template>
	<!-- Add more columns like this <ng-template #columnContent> Another column </ng-template> -->
	<ng-template #mainContent> Main content here </ng-template>
</app-page>
```

### Two columns - icons + column + content

```html
<app-page
	title="Title"
	[columnWidths]="['3rem', '25%']"
	[mobileColumnWidths]="['3rem', 'calc(100% - 3rem)']"
>
	<ng-template #columnContent>
		<div>
			<app-icon-menu></app-icon-menu>
		</div>
	</ng-template>
	<!-- Add more columns like this <ng-template #columnContent> Another column </ng-template> -->
	<ng-template #mainContent> Main content here </ng-template>
</app-page>
```
