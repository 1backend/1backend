"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7385],{80677:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>b,contentTitle:()=>f,default:()=>J,frontMatter:()=>k,metadata:()=>n,toc:()=>M});const n=JSON.parse('{"id":"1backend/make-default","title":"Make a Model Default","description":"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.","source":"@site/docs/1backend/make-default.api.mdx","sourceDirName":"1backend","slug":"/1backend/make-default","permalink":"/docs/1backend/make-default","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"make-default","title":"Make a Model Default","description":"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.","sidebar_label":"Make a Model Default","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VU1v4zYQ/SvEnFqAlpykC7Q61YstCnd3m6DenLI+jKWxxTVFcknKrlcg0B/RX9hfUozkryROgV5ysSVxZvje8L1hBxWF0isXlTVQwIxiECgaW5EWGESsSVS0xFbH/cd//vpbbGsywnnbOI72JAKZKLYq1raNAsXHPnL6Tl7IV0G0gaoMJFhHHnnjaQUFNLimd0MoSHDosaFIPkDx0D1BeagPEhS/O4w1SDDYEBfi1WkFEjx9bZWnCoroW5IQypoahKKDuHMcGqJXZgUpzTk4OGsCBV6/Ho/57/G2t+9BQmlNJBN5FZ3TquwZ5F8Ch3TPt7CLL1RGSCklCT9cKjs1G9SqEr/Nbn//Pxs4zw2MakBM3lt/iZp8AcnVcyT3BttYW6++UfVqSN5c7kkkb1CLGfkNefFLX/N1ICUJgcrWq7jrtfeW0JOftLGG4mHOUom4YlnudTjblDCX0FCsLQvZtYOAOR7yXo2jsCmHp7zbyzPlLPhRdVR86JkOcm+95txc2xJ1bUMs3vx4fXMFvPkB24y5DvTOET7t5KedI/F5H/IZxNJqbbdUicVOoAgOSxJoKhHtmozAcnCMWHrb9O69D+SZovhgV8oIMpWzysTs4L2asCJ/ct9kr6D+YODYX3TqPe36jiuztIyTzxLL/iypQcWMA2oKPwdlVq3G6K3JStuc1b6bilnrnPXcsKFJdYyuyPOrBZZrMhUn5JDkky5MRGvUUjHtIU5UtCFtXcOTy2mMS+sbsbReNKr0lg9DlRRGCwxUiclUnMktMHmtSjKBGP0B3a93H8TmJhs/whaKPN9ut9nKtJn1q3yfF3JcOT26ycZZHRvNiCP5JtwuZ8PWJ2phi6sV+UzZvA/Juasqag65ejvQAQksnoHsOLvJxiNfZtc/cV1nQ2zQnCH9iGs6junTzH3Usu7ktde9FfaCifRnzJ1GZZhD389ub6oHOJoK5PAMEorT3H/krLkEdhCndR0f573XKfHnry15tvhcwga9wgW3lC8bFfi5gmKJOtB/NOa7P/Y3zPfi7E66yODgA7Pjs0Ld8htIWNPu7M5K8yQPlmIow+qkLMnFs7xnsy+dD6C7+08gAffj4GRALiYPD1z9IqanBh4g8G+SL6R03WDvlI7xw9KLGcepMURzi+YppX8BSOPuHg==","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get a Model","permalink":"/docs/1backend/get-model"},"next":{"title":"Start a Model","permalink":"/docs/1backend/start-model"}}');var s=a(74848),o=a(28453),d=a(53746),i=a.n(d),l=a(56518),r=a.n(l),c=a(99972),p=a.n(c),u=a(25342),h=a.n(u),m=(a(44215),a(82223),a(24861));const k={id:"make-default",title:"Make a Model Default",description:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.",sidebar_label:"Make a Model Default",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU1v4zYQ/SvEnFqAlpykC7Q61YstCnd3m6DenLI+jKWxxTVFcknKrlcg0B/RX9hfUozkryROgV5ysSVxZvje8L1hBxWF0isXlTVQwIxiECgaW5EWGESsSVS0xFbH/cd//vpbbGsywnnbOI72JAKZKLYq1raNAsXHPnL6Tl7IV0G0gaoMJFhHHnnjaQUFNLimd0MoSHDosaFIPkDx0D1BeagPEhS/O4w1SDDYEBfi1WkFEjx9bZWnCoroW5IQypoahKKDuHMcGqJXZgUpzTk4OGsCBV6/Ho/57/G2t+9BQmlNJBN5FZ3TquwZ5F8Ch3TPt7CLL1RGSCklCT9cKjs1G9SqEr/Nbn//Pxs4zw2MakBM3lt/iZp8AcnVcyT3BttYW6++UfVqSN5c7kkkb1CLGfkNefFLX/N1ICUJgcrWq7jrtfeW0JOftLGG4mHOUom4YlnudTjblDCX0FCsLQvZtYOAOR7yXo2jsCmHp7zbyzPlLPhRdVR86JkOcm+95txc2xJ1bUMs3vx4fXMFvPkB24y5DvTOET7t5KedI/F5H/IZxNJqbbdUicVOoAgOSxJoKhHtmozAcnCMWHrb9O69D+SZovhgV8oIMpWzysTs4L2asCJ/ct9kr6D+YODYX3TqPe36jiuztIyTzxLL/iypQcWMA2oKPwdlVq3G6K3JStuc1b6bilnrnPXcsKFJdYyuyPOrBZZrMhUn5JDkky5MRGvUUjHtIU5UtCFtXcOTy2mMS+sbsbReNKr0lg9DlRRGCwxUiclUnMktMHmtSjKBGP0B3a93H8TmJhs/whaKPN9ut9nKtJn1q3yfF3JcOT26ycZZHRvNiCP5JtwuZ8PWJ2phi6sV+UzZvA/Juasqag65ejvQAQksnoHsOLvJxiNfZtc/cV1nQ2zQnCH9iGs6junTzH3Usu7ktde9FfaCifRnzJ1GZZhD389ub6oHOJoK5PAMEorT3H/krLkEdhCndR0f573XKfHnry15tvhcwga9wgW3lC8bFfi5gmKJOtB/NOa7P/Y3zPfi7E66yODgA7Pjs0Ld8htIWNPu7M5K8yQPlmIow+qkLMnFs7xnsy+dD6C7+08gAffj4GRALiYPD1z9IqanBh4g8G+SL6R03WDvlI7xw9KLGcepMURzi+YppX8BSOPuHg==",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,b={},M=[];function y(e){const t={p:"p",...(0,o.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Make a Model Default"}),"\n",(0,s.jsx)(i(),{method:"put",path:"/model-svc/model/{modelId}/make-default",context:"endpoint"}),"\n",(0,s.jsx)(t.p,{children:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used."}),"\n",(0,s.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,s.jsx)(r(),{parameters:[{description:"Model ID",in:"path",name:"modelId",required:!0,schema:{type:"string"}}]}),"\n",(0,s.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,s.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function J(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(y,{...e})}):y(e)}}}]);