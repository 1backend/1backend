"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1519],{82883:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>b,contentTitle:()=>m,default:()=>x,frontMatter:()=>u,metadata:()=>i,toc:()=>y});const i=JSON.parse('{"id":"1backend/get-config","title":"Get Config","description":"Fetch the current configuration from the server","source":"@site/docs/1backend/get-config.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-config","permalink":"/docs/1backend/get-config","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-config","title":"Get Config","description":"Fetch the current configuration from the server","sidebar_label":"Get Config","hide_title":true,"hide_table_of_contents":true,"api":"eJylVc1u4zYQfhWCZ0Vyki7Q6lTvYhukXbQBvDllfaCpscQNRXKHlF1VELAP0SfskxRDyrFiewu0vUgC5++bb76hBl6Bl6hcUNbwkv8EQTYsNMBkhwgmMGnNVtUdCvJgW7RtNHvAHSDPuHWQbPcVL3kN4V0M4Bl3AkULAdDz8mk4KfSraME7IYFnXNHBlw6w5xk3ogVextfB7mUDreDlwEPvyOgDKlPzcVxnHME7azx4st8sFvR6XerdxU4QAirYQcV8JyV4v+20JgDSmgAmUB7hnFYy+hefPSUbZmAcUu9BpdIp+/l5JUL0FlWlKJHQDzN7wA6yQ1t28xlk4GMWg36eCp70nHFVXTw+MnaBKOLpS6cQKl4+JUzrs7rj+Qmdfbe4Pif10YguNBbVH1D9G9JOgVGBN5emdm8CoBGaraLS2HtEi/+v0phxD7JDFfqoyLcgEHDZhYaXT2sSUxA1iZUnDbPVThJNLYTGTuqOuqYAXqSJX/mdnD5JqhFsEnyHmtwKbaXQjfWhfPP9ze01p0IHHCuCm4QyR3NKxsfeAfs0uXzibGu1tnuo2KZngsWpM2EqFuwzGCZkGvVxWx89IHXDPthaGQamclaZkB+WrwFRxW2etm85zTZyy19kIZz6BfooFGW2dpJ9EDKOA1qhqGMvNPgfvTJ1p0VAa3Jp21nuh3u26pyzSGQmkpoQXFkU1xshn8FUFFDEPXjFwpJJ27bWsMmNbS2y3nbIlvdsJgb/19c/WyXR0jSUBH+1ER6qyNCmUzqwYJmXQgMRoJUE4+PWHBDePXxgu9t88QqfL4tiv9/ntelyi3UxxflC1E5f3eaLvAmtJtQBsPW/bVep+rE9vxd1DZgrW0SXgphVQZPL9dvUE884CSg1vMhv88UVyvzmB8rrrA+tMDOkdxDYy337iqzhuCj/4VKf5h3g91A4LZSh8pGKYZL/Ez/KP20lgVhnnIRO5mEg0h9RjyMdp+s9/geUFxtNF9FWaA//gHv+j7iI6Bn6k1/FTuiO/OKS7QQqKhW3OzuInDCkwKWU4MIs6uxCGefbf/f+I8+4mBb0uBKULDt8UPaDyfSz3KcrlSDQc8y+ETIMaeHG8cU/mb4Z8bLHyZtYWo/j+DdVG7hy","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get Threads","permalink":"/docs/1backend/get-threads"},"next":{"title":"Save Config","permalink":"/docs/1backend/save-config"}}');var s=n(74848),a=n(28453),r=n(53746),c=n.n(r),o=n(56518),d=n.n(o),l=n(99972),p=n.n(l),g=n(25342),h=n.n(g),f=(n(44215),n(82223),n(24861));const u={id:"get-config",title:"Get Config",description:"Fetch the current configuration from the server",sidebar_label:"Get Config",hide_title:!0,hide_table_of_contents:!0,api:"eJylVc1u4zYQfhWCZ0Vyki7Q6lTvYhukXbQBvDllfaCpscQNRXKHlF1VELAP0SfskxRDyrFiewu0vUgC5++bb76hBl6Bl6hcUNbwkv8EQTYsNMBkhwgmMGnNVtUdCvJgW7RtNHvAHSDPuHWQbPcVL3kN4V0M4Bl3AkULAdDz8mk4KfSraME7IYFnXNHBlw6w5xk3ogVextfB7mUDreDlwEPvyOgDKlPzcVxnHME7azx4st8sFvR6XerdxU4QAirYQcV8JyV4v+20JgDSmgAmUB7hnFYy+hefPSUbZmAcUu9BpdIp+/l5JUL0FlWlKJHQDzN7wA6yQ1t28xlk4GMWg36eCp70nHFVXTw+MnaBKOLpS6cQKl4+JUzrs7rj+Qmdfbe4Pif10YguNBbVH1D9G9JOgVGBN5emdm8CoBGaraLS2HtEi/+v0phxD7JDFfqoyLcgEHDZhYaXT2sSUxA1iZUnDbPVThJNLYTGTuqOuqYAXqSJX/mdnD5JqhFsEnyHmtwKbaXQjfWhfPP9ze01p0IHHCuCm4QyR3NKxsfeAfs0uXzibGu1tnuo2KZngsWpM2EqFuwzGCZkGvVxWx89IHXDPthaGQamclaZkB+WrwFRxW2etm85zTZyy19kIZz6BfooFGW2dpJ9EDKOA1qhqGMvNPgfvTJ1p0VAa3Jp21nuh3u26pyzSGQmkpoQXFkU1xshn8FUFFDEPXjFwpJJ27bWsMmNbS2y3nbIlvdsJgb/19c/WyXR0jSUBH+1ER6qyNCmUzqwYJmXQgMRoJUE4+PWHBDePXxgu9t88QqfL4tiv9/ntelyi3UxxflC1E5f3eaLvAmtJtQBsPW/bVep+rE9vxd1DZgrW0SXgphVQZPL9dvUE884CSg1vMhv88UVyvzmB8rrrA+tMDOkdxDYy337iqzhuCj/4VKf5h3g91A4LZSh8pGKYZL/Ez/KP20lgVhnnIRO5mEg0h9RjyMdp+s9/geUFxtNF9FWaA//gHv+j7iI6Bn6k1/FTuiO/OKS7QQqKhW3OzuInDCkwKWU4MIs6uxCGefbf/f+I8+4mBb0uBKULDt8UPaDyfSz3KcrlSDQc8y+ETIMaeHG8cU/mb4Z8bLHyZtYWo/j+DdVG7hy",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,b={},y=[];function v(e){const t={p:"p",...(0,a.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"Get Config"}),"\n",(0,s.jsx)(c(),{method:"get",path:"/config-svc/config",context:"endpoint"}),"\n",(0,s.jsx)(t.p,{children:"Fetch the current configuration from the server"}),"\n",(0,s.jsx)(f.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,s.jsx)(d(),{parameters:[{description:"Namespace",in:"query",name:"namespace",schema:{type:"string"}}]}),"\n",(0,s.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,s.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"Current configuration retrieved successfully",content:{"application/json":{schema:{properties:{config:{properties:{data:{additionalProperties:!0,type:"object"},dataJson:{type:"string"},id:{type:"string"},namespace:{type:"string"}},required:["data"],type:"object"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function x(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(v,{...e})}):v(e)}}}]);