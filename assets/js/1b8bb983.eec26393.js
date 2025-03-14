"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5424],{76373:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>f,contentTitle:()=>g,default:()=>E,frontMatter:()=>m,metadata:()=>a,toc:()=>b});const a=JSON.parse('{"id":"1backend/get-thread","title":"Get Thread","description":"Fetch information about a specific chat thread by its ID","source":"@site/docs/1backend/get-thread.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-thread","permalink":"/docs/1backend/get-thread","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-thread","title":"Get Thread","description":"Fetch information about a specific chat thread by its ID","sidebar_label":"Get Thread","hide_title":true,"hide_table_of_contents":true,"api":"eJylVt1u2zYUfhWCVxugSE6yApuu5m5d4K5Yuzm9SoyBpo4lNhTJ8hzZ9QwBfYg+4Z5kOJLsKIlarOiNLZHn5zsfz/mogywAdTSBjHcyl78B6UoYt/GxVrwm1No3JJTAANpsjBa6UiSoiqAKsd4LQygWv8pE+gCxc1kUMpcl0HVnIxMZVFQ1EESU+c3hUcbeqg9heCEoqmQinapB5rJPtOAwEd43JkIhc4oNJBJ1BbWS+UHSPrAtUjSulG27YmMM3iEg71/MZvw3mbgAUsaiwEZrQNw01u5FBIoGtsBptXcEjjiACsEa3RWZvUOOchihCJEpINPnhA8GCUfo1t5bUE62yVDUUx8dQREUc5ooKpGm6OOqOtiBmr+hXv7pwst1dSWTpx5kiC2fVM7Lwm8EVTCcZDrp7oPRiwInIgw7ooCNcYBiVxldic4BR2HFGqx3JQry6a277re1ciL6pqzsXqzZkp+J8SgUpEoUGx+HAMjADEGNk5QMCypGtef3JhRfILBBiJPlvO03OuCmwCM1bM+l+Q4yAmM1OGLs/wJrx817wwe5Opn49TvQ1Nk8XuG1H6Zad+G2yppCvFy+/uNrOvTxnPQJzicIcaqhykfzz9eNwFSCZ9MVEESnrFhC3EIUL2L08dsytYlE0E00tO9k5jmoCHHeUCXzmxVrAvcW8/8LK9hyq/kUaqDKD4LVSRWby4xF7gy3OusPOzscdaiVnIYx92LWRMv2mfVa2coj5c9+vLg8l5zvCGfJqPsJH4N6MlT7AOJ2MLmVYuOt9TvoVJb1V2kQyhWC/B04oXTfUGITfd21K3cxlyVe+dI4Aa4I3jhKj8JagSog3kvrfDjijuL7+VfB/A590/I90OmSd6R0dypQK8MVo7KAP6NxZWMVRe9S7etR7DcLsWxC8JFZ7UmqiEKeZedrpe/AFeyQ8aw8ZGEutK9r78Rg1onB3jdRzBdi1BP478dPtdHR82kYDXi2VghFx9C6MZYEeYFaWWACrNHgsNPCI8KrN6/E9jKdPcCHeZbtdru0dE3qY5kNfpipMtizy3SWVlTbbsIh1vh6s+yz35eHO1WWEFPjs84kkycdlufP+5pkIrmB+oJn6WU6O4s6vfiJ4waPVCs3QnoFJE536QOyDvfz8i0393DwBB8oC1aZ7o7qODkMA3EjjwMhT7dXIvPT5bxKJPc+Gx4OfA5vo21bXn7fQOR5XCVyq6JRayaCPwIM8nMh842yCF8o7Lu/Bun8Xoy/FSZBH1vY7ZliZRt+k4m8g/34W6JdtclxHBhMvz3XGgKNHJ8oUDsWjKsX1zKRahjl++HhYMnxgaNPgno8fD0E/m2Tz7gcDv1otu3Jvt/6rMdp4ntr5mjVtu1/jdVqfg==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Delete a Thread","permalink":"/docs/1backend/delete-thread"},"next":{"title":"Update Thread","permalink":"/docs/1backend/update-thread"}}');var s=i(74848),d=i(28453),r=i(53746),o=i.n(r),n=i(56518),c=i.n(n),p=i(99972),h=i.n(p),l=i(25342),u=i.n(l),y=(i(44215),i(82223),i(24861));const m={id:"get-thread",title:"Get Thread",description:"Fetch information about a specific chat thread by its ID",sidebar_label:"Get Thread",hide_title:!0,hide_table_of_contents:!0,api:"eJylVt1u2zYUfhWCVxugSE6yApuu5m5d4K5Yuzm9SoyBpo4lNhTJ8hzZ9QwBfYg+4Z5kOJLsKIlarOiNLZHn5zsfz/mogywAdTSBjHcyl78B6UoYt/GxVrwm1No3JJTAANpsjBa6UiSoiqAKsd4LQygWv8pE+gCxc1kUMpcl0HVnIxMZVFQ1EESU+c3hUcbeqg9heCEoqmQinapB5rJPtOAwEd43JkIhc4oNJBJ1BbWS+UHSPrAtUjSulG27YmMM3iEg71/MZvw3mbgAUsaiwEZrQNw01u5FBIoGtsBptXcEjjiACsEa3RWZvUOOchihCJEpINPnhA8GCUfo1t5bUE62yVDUUx8dQREUc5ooKpGm6OOqOtiBmr+hXv7pwst1dSWTpx5kiC2fVM7Lwm8EVTCcZDrp7oPRiwInIgw7ooCNcYBiVxldic4BR2HFGqx3JQry6a277re1ciL6pqzsXqzZkp+J8SgUpEoUGx+HAMjADEGNk5QMCypGtef3JhRfILBBiJPlvO03OuCmwCM1bM+l+Q4yAmM1OGLs/wJrx817wwe5Opn49TvQ1Nk8XuG1H6Zad+G2yppCvFy+/uNrOvTxnPQJzicIcaqhykfzz9eNwFSCZ9MVEESnrFhC3EIUL2L08dsytYlE0E00tO9k5jmoCHHeUCXzmxVrAvcW8/8LK9hyq/kUaqDKD4LVSRWby4xF7gy3OusPOzscdaiVnIYx92LWRMv2mfVa2coj5c9+vLg8l5zvCGfJqPsJH4N6MlT7AOJ2MLmVYuOt9TvoVJb1V2kQyhWC/B04oXTfUGITfd21K3cxlyVe+dI4Aa4I3jhKj8JagSog3kvrfDjijuL7+VfB/A590/I90OmSd6R0dypQK8MVo7KAP6NxZWMVRe9S7etR7DcLsWxC8JFZ7UmqiEKeZedrpe/AFeyQ8aw8ZGEutK9r78Rg1onB3jdRzBdi1BP478dPtdHR82kYDXi2VghFx9C6MZYEeYFaWWACrNHgsNPCI8KrN6/E9jKdPcCHeZbtdru0dE3qY5kNfpipMtizy3SWVlTbbsIh1vh6s+yz35eHO1WWEFPjs84kkycdlufP+5pkIrmB+oJn6WU6O4s6vfiJ4waPVCs3QnoFJE536QOyDvfz8i0393DwBB8oC1aZ7o7qODkMA3EjjwMhT7dXIvPT5bxKJPc+Gx4OfA5vo21bXn7fQOR5XCVyq6JRayaCPwIM8nMh842yCF8o7Lu/Bun8Xoy/FSZBH1vY7ZliZRt+k4m8g/34W6JdtclxHBhMvz3XGgKNHJ8oUDsWjKsX1zKRahjl++HhYMnxgaNPgno8fD0E/m2Tz7gcDv1otu3Jvt/6rMdp4ntr5mjVtu1/jdVqfg==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},g=void 0,f={},b=[];function J(e){const t={p:"p",...(0,d.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(y.default,{as:"h1",className:"openapi__heading",children:"Get Thread"}),"\n",(0,s.jsx)(o(),{method:"get",path:"/chat-svc/thread/{threadId}",context:"endpoint"}),"\n",(0,s.jsx)(t.p,{children:"Fetch information about a specific chat thread by its ID"}),"\n",(0,s.jsx)(y.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,s.jsx)(c(),{parameters:[{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}]}),"\n",(0,s.jsx)(h(),{title:"Body",body:void 0}),"\n",(0,s.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"Thread details successfully retrieved",content:{"application/json":{schema:{properties:{exists:{type:"boolean"},thread:{properties:{createdAt:{type:"string"},id:{example:"thr_emSQnpJbhG",type:"string"},title:{description:"Title of the thread.",type:"string"},topicIds:{description:"TopicIds defines which topics the thread belongs to.\nTopics can roughly be thought of as tags for threads.",items:{type:"string"},type:"array"},updatedAt:{type:"string"},userIds:{description:"UserIds the ids of the users who can see this thread.",items:{type:"string"},type:"array"}},required:["id"],type:"object"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function E(e={}){const{wrapper:t}={...(0,d.R)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(J,{...e})}):J(e)}}}]);