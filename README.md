# 📊 Backend Skills Coverage Tracker

<style>
* { box-sizing: border-box; }
.wrap { padding: 1rem; }
.category { margin-bottom: 1rem; border-radius: 8px; border: 0.5px solid #ddd; overflow: hidden; }
.cat-header { display: flex; align-items: center; gap: 10px; padding: 10px 14px; cursor: pointer; user-select: none; }
.cat-title { font-size: 13px; font-weight: 500; }
.cat-count { font-size: 11px; padding: 2px 7px; border-radius: 10px; font-weight: 500; }
.chevron { margin-left: auto; font-size: 11px; transition: transform 0.2s; }
.chevron.open { transform: rotate(180deg); }
.cat-body { display: block; padding: 0 14px 12px; }
.pills { display: flex; flex-wrap: wrap; gap: 6px; margin-top: 6px; }
.pill { font-size: 12px; padding: 3px 10px; border-radius: 20px; cursor: pointer; border: 0.5px solid #ccc; background: #f5f5f5; }
.pill.have { background: #EAF3DE; color: #27500A; border-color: #97C459; }
.progress-bar-wrap { height: 4px; background: #eee; border-radius: 2px; margin-top: 6px; }
.progress-bar { height: 100%; border-radius: 2px; }
.legend { display: flex; gap: 16px; margin-bottom: 12px; font-size: 12px; align-items: center; flex-wrap: wrap; }
.legend-dot { width: 10px; height: 10px; border-radius: 50%; display: inline-block; margin-right: 4px; }
.summary { display: grid; grid-template-columns: repeat(3, 1fr); gap: 8px; margin-bottom: 12px; }
.stat { background: #f5f5f5; border-radius: 6px; padding: 8px 10px; text-align: center; }
.stat-val { font-size: 18px; font-weight: 500; }
.stat-label { font-size: 11px; margin-top: 2px; }
</style>

<div class="wrap">

<div class="legend">
  <span><span class="legend-dot" style="background:#97C459"></span>You have this</span>
  <span><span class="legend-dot" style="background:#ccc"></span>To learn</span>
</div>

<div class="summary">
  <div class="stat">
    <div class="stat-val">44</div>
    <div class="stat-label">You have</div>
  </div>
  <div class="stat">
    <div class="stat-val">102</div>
    <div class="stat-label">Total topics</div>
  </div>
  <div class="stat">
    <div class="stat-val">43%</div>
    <div class="stat-label">Coverage</div>
  </div>
</div>

---

## 🧩 API Design
<div class="category">
<div class="cat-header">
<span>🔵</span><span class="cat-title">API Design</span><span class="cat-count">6/11</span>
</div>
<div class="cat-body">

<div class="progress-bar-wrap">
<div class="progress-bar" style="width:54%;background:#378ADD"></div>
</div>

<div class="pills">
<span class="pill have">REST API</span>
<span class="pill have">gRPC + Protobuf</span>
<span class="pill have">GraphQL</span>
<span class="pill">WebSockets</span>
<span class="pill">API versioning</span>
<span class="pill">Rate limiting</span>
<span class="pill have">API Gateway</span>
<span class="pill have">OpenAPI / Swagger</span>
<span class="pill">Idempotency</span>
<span class="pill">Pagination patterns</span>
<span class="pill">HATEOAS</span>
</div>

</div>
</div>

---

## 🏗️ Architecture Patterns
<div class="category">
<div class="cat-body">

<div class="progress-bar-wrap">
<div class="progress-bar" style="width:33%;background:#7F77DD"></div>
</div>

<div class="pills">
<span class="pill have">Microservices</span>
<span class="pill have">Monolithic</span>
<span class="pill have">Distributed systems</span>
<span class="pill have">Event-driven (EDA)</span>
<span class="pill">CQRS</span>
<span class="pill">Event sourcing</span>
<span class="pill">Saga pattern</span>
<span class="pill">Hexagonal architecture</span>
<span class="pill">Domain-driven design</span>
<span class="pill">Service mesh</span>
<span class="pill">BFF pattern</span>
<span class="pill">Strangler fig</span>
</div>

</div>
</div>

---

## 🗄️ Databases & Storage
<div class="category">
<div class="cat-body">

<div class="progress-bar-wrap">
<div class="progress-bar" style="width:47%;background:#1D9E75"></div>
</div>

<div class="pills">
<span class="pill have">PostgreSQL</span>
<span class="pill have">MySQL</span>
<span class="pill have">MongoDB</span>
<span class="pill have">Redis</span>
<span class="pill have">Cassandra</span>
<span class="pill have">Database indexing</span>
<span class="pill have">Query optimisation</span>
<span class="pill">ACID vs BASE</span>
<span class="pill">CAP theorem</span>
<span class="pill">Database sharding</span>
<span class="pill">Replication</span>
<span class="pill">Connection pooling</span>
<span class="pill">Data partitioning</span>
<span class="pill">Time-series DBs</span>
<span class="pill">Elasticsearch</span>
</div>

</div>
</div>

---

## 🔐 Security
<div class="category">
<div class="cat-body">

<div class="progress-bar-wrap">
<div class="progress-bar" style="width:36%;background:#E24B4A"></div>
</div>

<div class="pills">
<span class="pill have">OAuth2.0</span>
<span class="pill have">OIDC</span>
<span class="pill have">SSO</span>
<span class="pill have">JWT</span>
<span class="pill">mTLS</span>
<span class="pill">Secret management</span>
<span class="pill">RBAC / ABAC</span>
<span class="pill">SQL injection prevention</span>
<span class="pill">OWASP top 10</span>
<span class="pill">Zero trust architecture</span>
<span class="pill">Encryption</span>
</div>

</div>
</div>

---

## 🧪 Testing & Quality
<div class="category">
<div class="cat-body">

<div class="progress-bar-wrap">
<div class="progress-bar" style="width:55%;background:#1D9E75"></div>
</div>

<div class="pills">
<span class="pill have">Unit testing</span>
<span class="pill have">Integration testing</span>
<span class="pill have">Load testing</span>
<span class="pill have">TDD</span>
<span class="pill">Contract testing</span>
<span class="pill">Mutation testing</span>
<span class="pill">Chaos testing</span>
<span class="pill">E2E testing</span>
<span class="pill have">SonarQube</span>
</div>

</div>
</div>

---

## 🤖 GenAI & ML Backend
<div class="category">
<div class="cat-body">

<div class="progress-bar-wrap">
<div class="progress-bar" style="width:44%;background:#D85A30"></div>
</div>

<div class="pills">
<span class="pill have">LLMs</span>
<span class="pill have">RAG pipelines</span>
<span class="pill have">AI Agents</span>
<span class="pill have">Fine-tuning</span>
<span class="pill">Vector databases</span>
<span class="pill">Embeddings</span>
<span class="pill">LLM orchestration</span>
<span class="pill">Prompt engineering</span>
<span class="pill">Model serving</span>
</div>

</div>
</div>

---

## 📌 Notes
- Green pills → already covered
- Grey pills → learning backlog
- Progress bars show category coverage
