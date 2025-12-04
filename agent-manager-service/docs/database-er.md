```mermaid
erDiagram
	ORGANIZATIONS {
		uuid org_id PK "NOT NULL" 
		text open_choreo_org UK "NOT NULL"  
		timestamptz created_at  "NOT NULL, DEFAULT CURRENT_TIMESTAMP"  
	}
```
