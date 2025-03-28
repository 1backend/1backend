"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7385],{80677:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>y,contentTitle:()=>k,default:()=>b,frontMatter:()=>h,metadata:()=>d,toc:()=>j});const d=JSON.parse('{"id":"1backend/make-default","title":"Make a Model Default","description":"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.","source":"@site/docs/1backend/make-default.api.mdx","sourceDirName":"1backend","slug":"/1backend/make-default","permalink":"/docs/1backend/make-default","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"make-default","title":"Make a Model Default","description":"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.","sidebar_label":"Make a Model Default","hide_title":true,"hide_table_of_contents":true,"api":"eJy9Vctu4zYU/RXirlqAkZxxByi0agZTFO7MNEE9WWW8uJauLY4pkkNSdl2BQD+iX9gvKa4kP5I4BbrJxpbE+zj38Byyg4pC6ZWLyhooYE4xCBSNrUgLDCLWJCpaYavj+PGfv/4Wu5qMcN42jqM9iUAmip2KtW2jQPGpj5y9lxfyVRBtoCoDCdaRR248q6CABjf0fggFCQ49NhTJBygeuicoD/VBguJ3h7EGCQYb4kK8OqtAgqdvrfJUQRF9SxJCWVODUHQQ945DQ/TKrCGlBQcHZ02gwOtvJhP+e9z29gNIKK2JZCKvonNalf0E+dfAId3zFnb5lcoIKaUk4YdLZWdmi1pV4tf57W//p4HzTGBUA2Ly3vpLo8kXkFw/R3JvsI219epPql4NydvLnETyBrWYk9+SFz/3NV8HUpIQqGy9ivtee+8IPfmbNtZQPCxYKhHXLMtRh/NtCQsJDcXaspBdOwiY4yHv1XgVtuXwlHejPFPOgr+qjooP/aSD3FuvOTfXtkRd2xCLtz++mV4DNz9gm/Osw3jnCJ8y+XnvSHwZQ76AWFmt7Y4qsdwLFMFhSQJNJaLdkBFYDo4RK2+b3r33gTyPKD7atTKCTOWsMjE7eK8mrMif3HczKqjfGDjyi059oH3PuDIryzh5L7Hs95IaVDxxQE3hp6DMutUYvTVZaZuz2nczMW+ds54JG0iqY3RFnl8vsdyQqTghhySfsHAzuzIY1ZZEo0pvmWtVUhBOY1xZ3/A4WpVkAjGeQ79f7j6K7TSbPOoWijzf7XbZ2rSZ9et8zAs5rp2+mmaTrI6NZgyRfBNuV/Oh2wls2OF6TT5TNu9DcuZJRc0h1++GQUACy2GAP8mm2eTKl9l0wnWdDbFBc4b0E27oePCeTtFHJHQn97zuOT9KINIfMXcaleEZej670SYPcLQJyOEZJBSnk/yRVxYS2BOc1nVLDHTvdUr8+VtLnk27kLBFr3DJlPL1oQI/V1CsUAf6D2K++328M74XZ7fMxQkOyjZ73ivULb+BhA3tz26htEjyYBKGMqzelCW5eJb37DRL50fK3f1nkICjwU+W4mLy8MDVL2J6askBAv8m+UJK1w2GTekYPyy9mHE8B4ZopmiRUvoXGm7dAA==","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get a Model","permalink":"/docs/1backend/get-model"},"next":{"title":"Start a Model","permalink":"/docs/1backend/start-model"}}');var o=a(74848),s=a(28453),n=a(53746),i=a.n(n),l=a(56518),r=a.n(l),p=a(99972),c=a.n(p),u=a(25342),m=a.n(u),f=(a(44215),a(82223),a(24861));const h={id:"make-default",title:"Make a Model Default",description:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.",sidebar_label:"Make a Model Default",hide_title:!0,hide_table_of_contents:!0,api:"eJy9Vctu4zYU/RXirlqAkZxxByi0agZTFO7MNEE9WWW8uJauLY4pkkNSdl2BQD+iX9gvKa4kP5I4BbrJxpbE+zj38Byyg4pC6ZWLyhooYE4xCBSNrUgLDCLWJCpaYavj+PGfv/4Wu5qMcN42jqM9iUAmip2KtW2jQPGpj5y9lxfyVRBtoCoDCdaRR248q6CABjf0fggFCQ49NhTJBygeuicoD/VBguJ3h7EGCQYb4kK8OqtAgqdvrfJUQRF9SxJCWVODUHQQ945DQ/TKrCGlBQcHZ02gwOtvJhP+e9z29gNIKK2JZCKvonNalf0E+dfAId3zFnb5lcoIKaUk4YdLZWdmi1pV4tf57W//p4HzTGBUA2Ly3vpLo8kXkFw/R3JvsI219epPql4NydvLnETyBrWYk9+SFz/3NV8HUpIQqGy9ivtee+8IPfmbNtZQPCxYKhHXLMtRh/NtCQsJDcXaspBdOwiY4yHv1XgVtuXwlHejPFPOgr+qjooP/aSD3FuvOTfXtkRd2xCLtz++mV4DNz9gm/Osw3jnCJ8y+XnvSHwZQ76AWFmt7Y4qsdwLFMFhSQJNJaLdkBFYDo4RK2+b3r33gTyPKD7atTKCTOWsMjE7eK8mrMif3HczKqjfGDjyi059oH3PuDIryzh5L7Hs95IaVDxxQE3hp6DMutUYvTVZaZuz2nczMW+ds54JG0iqY3RFnl8vsdyQqTghhySfsHAzuzIY1ZZEo0pvmWtVUhBOY1xZ3/A4WpVkAjGeQ79f7j6K7TSbPOoWijzf7XbZ2rSZ9et8zAs5rp2+mmaTrI6NZgyRfBNuV/Oh2wls2OF6TT5TNu9DcuZJRc0h1++GQUACy2GAP8mm2eTKl9l0wnWdDbFBc4b0E27oePCeTtFHJHQn97zuOT9KINIfMXcaleEZej670SYPcLQJyOEZJBSnk/yRVxYS2BOc1nVLDHTvdUr8+VtLnk27kLBFr3DJlPL1oQI/V1CsUAf6D2K++328M74XZ7fMxQkOyjZ73ivULb+BhA3tz26htEjyYBKGMqzelCW5eJb37DRL50fK3f1nkICjwU+W4mLy8MDVL2J6askBAv8m+UJK1w2GTekYPyy9mHE8B4ZopmiRUvoXGm7dAA==",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},k=void 0,y={},j=[];function D(e){const t={p:"p",...(0,s.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"Make a Model Default"}),"\n",(0,o.jsx)(i(),{method:"put",path:"/model-svc/model/{modelId}/make-default",context:"endpoint"}),"\n",(0,o.jsx)(t.p,{children:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used."}),"\n",(0,o.jsx)(f.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,o.jsx)(r(),{parameters:[{description:"Model ID",in:"path",name:"modelId",required:!0,schema:{type:"string"}}]}),"\n",(0,o.jsx)(c(),{title:"Body",body:void 0}),"\n",(0,o.jsx)(m(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function b(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,o.jsx)(t,{...e,children:(0,o.jsx)(D,{...e})}):D(e)}}}]);