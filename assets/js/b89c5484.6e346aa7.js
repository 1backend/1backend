"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2834],{83594:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>j,contentTitle:()=>y,default:()=>x,frontMatter:()=>v,metadata:()=>i,toc:()=>h});const i=JSON.parse('{"id":"1backend/remove-instance","title":"Remove Instance","description":"Removes a registered instance by ID.","source":"@site/docs/1backend/remove-instance.api.mdx","sourceDirName":"1backend","slug":"/1backend/remove-instance","permalink":"/docs/1backend/remove-instance","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"remove-instance","title":"Remove Instance","description":"Removes a registered instance by ID.","sidebar_label":"Remove Instance","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VU1v2zgQ/SvEnHYBRXKaFljotGkTLIwNdoukOaU+0NRYYkOR7JCy1yvwvy9GH4mduIuecrH18Ubz3vA9socKgyLto3YWSrjF1m0xCCkIax0iElZC2xClVSjWe7G8yiED55EklywrKIGGouWEggy8JNliRApQPvQvWsw4sbyCDDQ/8jI2kIGVLUIJuoIMCL93mrCCMlKHGQTVYCuh7CHuPaNCJG1rSGnF4OCdDRj4/bvFe/47bvqXE5+cjWgjpAzeLxavIUu7lUZXIy01gcsepPdGq0Ft8S0wtD9g44lnEfXYG4kcnSKZzU/c+huqCCmlgcf5ax73VnaxcaT/xeoNmZwY2h3SVisU1kWxcZ19OzofTi9QRLLSCOaFJK6Hb74NpZRBQNWRjvvB0h9REtJlFxsoH1bswShrdjvcDrmhvbjbKlhl0GJsHKekQoNxTAdXQUET8ixsVTFnrOh1lYC7scYxPx0ZxhfGKWkaF2L54bd3F+fAbWdWd6xyFHbI7eUMv+w9iq8T5CuIjTPG7bDiaEsRvFQopK1EdI9ohVRjCMWGXCtig+I+ILEyceNqbQXayjttYz4nuUFZIT1n+XKy8rAk8DRZ6fWfuB9mre3GMU9eRamGVcRWalYcpMHwe9C27oyM5GyuXHvw7c9Lcdd57yhCNg2pidGXRXG+luoRbcUFBSf+eAqXyzMro96iaLUiF0afB+GNjBtHLcsxWqENyHzmfn98vhHbi3xx1C2URbHb7fLadrmjupjqQiFrb84u8kXexNYwh4jUhr83U6qeyYadrGukXLtigBQ8Jx0NQ84/jkIgA7bDSH+RX+SLM1I5eyAD70JspT1gOu7i4mBHPtLfP0fmZzf8ad0i/hMLb6S23HgYQj/5+QEO/Tz44al7qSuOAluXgX2/lgHvyaTEj793SJyqVQZbSVquWTkfGzrwdQXlRpqA/yPil9vptPhVHJ8uJ2nPHrR7nqo0Hd9BBo+4H0+ftErZ7GQmMr64VAp9PCh5tdmkw7hfXd9cf7mGDOQUxGfr8/ey+YIbnGT0MjojC/5N2Q9K+n4MVkpP+PHVDyue8jqJZxEppf8ArkjDAw==","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Register Instance","permalink":"/docs/1backend/register-instance"},"next":{"title":"List Service Instances","permalink":"/docs/1backend/list-instances"}}');var s=n(74848),a=n(28453),r=n(53746),o=n.n(r),c=n(56518),d=n.n(c),l=n(99972),p=n.n(l),u=n(25342),b=n.n(u),m=(n(44215),n(82223),n(24861));const v={id:"remove-instance",title:"Remove Instance",description:"Removes a registered instance by ID.",sidebar_label:"Remove Instance",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU1v2zgQ/SvEnHYBRXKaFljotGkTLIwNdoukOaU+0NRYYkOR7JCy1yvwvy9GH4mduIuecrH18Ubz3vA9socKgyLto3YWSrjF1m0xCCkIax0iElZC2xClVSjWe7G8yiED55EklywrKIGGouWEggy8JNliRApQPvQvWsw4sbyCDDQ/8jI2kIGVLUIJuoIMCL93mrCCMlKHGQTVYCuh7CHuPaNCJG1rSGnF4OCdDRj4/bvFe/47bvqXE5+cjWgjpAzeLxavIUu7lUZXIy01gcsepPdGq0Ft8S0wtD9g44lnEfXYG4kcnSKZzU/c+huqCCmlgcf5ax73VnaxcaT/xeoNmZwY2h3SVisU1kWxcZ19OzofTi9QRLLSCOaFJK6Hb74NpZRBQNWRjvvB0h9REtJlFxsoH1bswShrdjvcDrmhvbjbKlhl0GJsHKekQoNxTAdXQUET8ixsVTFnrOh1lYC7scYxPx0ZxhfGKWkaF2L54bd3F+fAbWdWd6xyFHbI7eUMv+w9iq8T5CuIjTPG7bDiaEsRvFQopK1EdI9ohVRjCMWGXCtig+I+ILEyceNqbQXayjttYz4nuUFZIT1n+XKy8rAk8DRZ6fWfuB9mre3GMU9eRamGVcRWalYcpMHwe9C27oyM5GyuXHvw7c9Lcdd57yhCNg2pidGXRXG+luoRbcUFBSf+eAqXyzMro96iaLUiF0afB+GNjBtHLcsxWqENyHzmfn98vhHbi3xx1C2URbHb7fLadrmjupjqQiFrb84u8kXexNYwh4jUhr83U6qeyYadrGukXLtigBQ8Jx0NQ84/jkIgA7bDSH+RX+SLM1I5eyAD70JspT1gOu7i4mBHPtLfP0fmZzf8ad0i/hMLb6S23HgYQj/5+QEO/Tz44al7qSuOAluXgX2/lgHvyaTEj793SJyqVQZbSVquWTkfGzrwdQXlRpqA/yPil9vptPhVHJ8uJ2nPHrR7nqo0Hd9BBo+4H0+ftErZ7GQmMr64VAp9PCh5tdmkw7hfXd9cf7mGDOQUxGfr8/ey+YIbnGT0MjojC/5N2Q9K+n4MVkpP+PHVDyue8jqJZxEppf8ArkjDAw==",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},y=void 0,j={},h=[];function f(e){const t={p:"p",...(0,a.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Remove Instance"}),"\n",(0,s.jsx)(o(),{method:"delete",path:"/registry-svc/instance/{id}",context:"endpoint"}),"\n",(0,s.jsx)(t.p,{children:"Removes a registered instance by ID."}),"\n",(0,s.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,s.jsx)(d(),{parameters:[{description:"Instance ID",in:"path",name:"id",required:!0,schema:{type:"string"}}]}),"\n",(0,s.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,s.jsx)(b(),{id:void 0,label:void 0,responses:{204:{description:"No Content"},400:{description:"Invalid ID",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},404:{description:"Service not found",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function x(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(f,{...e})}):f(e)}}}]);