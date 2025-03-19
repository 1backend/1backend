"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2603],{54566:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>I,contentTitle:()=>k,default:()=>h,frontMatter:()=>f,metadata:()=>a,toc:()=>v});const a=JSON.parse('{"id":"1backend/container-daemon-info","title":"Get Container Daemon Information","description":"Retrieve detailed information about the availability and status of container daemons on the node.","source":"@site/docs/1backend/container-daemon-info.api.mdx","sourceDirName":"1backend","slug":"/1backend/container-daemon-info","permalink":"/docs/1backend/container-daemon-info","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"container-daemon-info","title":"Get Container Daemon Information","description":"Retrieve detailed information about the availability and status of container daemons on the node.","sidebar_label":"Get Container Daemon Information","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VE1v2zgQ/SvEnBXJSbbAVqem3SIwWqDFpjm5PoypscyGIlmSsusV9N8XQ8pf2XQPPfRiC+LM6L3H92aAhoL0ykVlDdTwN0WvaEuioYhKUyOUWVvfIZ8LXNk+irghgVtUGldKq7gXaBoRIsY+CLsW0pqIypAXDVJnTRDWpB5jGyqhAOvIp3nzBmo4lv+VqudmbaEAT8FZEyhAPcDNbMZ/l0gfyG+VJDE/4YMiTSMTuRyd00qmg+pb4J4BgtxQh/zkPMOIKn8Bm8ZTSI9x7whqCNEr08JYwERV09npylpNaPiYvLf+hcaRSXzvlacG6sXZlGVxKLWrbyQjjCMX/zG7/i/JR4N93Fiv/qHm19n9D8SXkLx6Se65ieQNasG6kxfv08zfA2ksIJDsvYp7qBcDvCX05O/6uIF6sRxZUGwDq/zu6L2HrWSlO4obyzZrKUIBDrkHqqPprsJWVtmnlcrWC4lgSF/qvebySluJemNDrF/9eXN7DfzNA6QHpphZnQN7LuCXvSPxdSr5CmJttbY7asRqL1AEh5JSkKJ9IiNQZuuItbddSs9jyKzER9sqI8g0zioTOVCK528IG+IbMdixfneTcQ7JmGRFpz7QPgmd+NZDTqBMV0gdKmYcUFN4E5Rpe43RW1NK253N/jwXD71z1rOoWaRNjK6uqusVyicyDTdUnI9LFe5Eb9RaMe1cJxrakrauIxOF0xg5zWJtveiU9DbklIerFQZqxN1cnLksMHmtJJmQwnlAd//5o9jelrMLbKGuqt1uV7amL61vq6kvVNg6fXVbzspN7DQjjuS78Gk9LZgTtbDDtiVfKlulkopVVZEXA1y/zXSgADZPJjsrb8vZlZflzWue62yIHZozpPcUxcmyeQE+W2gX8g2nuP2WTT15JtKPWDmNKi28JOkwRWkBF1FiwGkUTPZaFsCx4cJh4Dt89Hoc+fX3njzHeVnAFr3KC3axHIuDkzl/T7TnC5WSHDtti7rPJn62acbzqN+//wIF4JTCk+95WHF44OmHI7M/m/08NxkC/47FT1qGIadqHI/1+einHcew5moWdjmO479zt51r","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"List Containers","permalink":"/docs/1backend/list-containers"},"next":{"title":"Get Container Host","permalink":"/docs/1backend/get-host"}}');var i=t(74848),o=t(28453),r=t(53746),s=t.n(r),d=t(56518),c=t.n(d),l=t(99972),p=t.n(l),m=t(25342),u=t.n(m),b=(t(44215),t(82223),t(24861));const f={id:"container-daemon-info",title:"Get Container Daemon Information",description:"Retrieve detailed information about the availability and status of container daemons on the node.",sidebar_label:"Get Container Daemon Information",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VE1v2zgQ/SvEnBXJSbbAVqem3SIwWqDFpjm5PoypscyGIlmSsusV9N8XQ8pf2XQPPfRiC+LM6L3H92aAhoL0ykVlDdTwN0WvaEuioYhKUyOUWVvfIZ8LXNk+irghgVtUGldKq7gXaBoRIsY+CLsW0pqIypAXDVJnTRDWpB5jGyqhAOvIp3nzBmo4lv+VqudmbaEAT8FZEyhAPcDNbMZ/l0gfyG+VJDE/4YMiTSMTuRyd00qmg+pb4J4BgtxQh/zkPMOIKn8Bm8ZTSI9x7whqCNEr08JYwERV09npylpNaPiYvLf+hcaRSXzvlacG6sXZlGVxKLWrbyQjjCMX/zG7/i/JR4N93Fiv/qHm19n9D8SXkLx6Se65ieQNasG6kxfv08zfA2ksIJDsvYp7qBcDvCX05O/6uIF6sRxZUGwDq/zu6L2HrWSlO4obyzZrKUIBDrkHqqPprsJWVtmnlcrWC4lgSF/qvebySluJemNDrF/9eXN7DfzNA6QHpphZnQN7LuCXvSPxdSr5CmJttbY7asRqL1AEh5JSkKJ9IiNQZuuItbddSs9jyKzER9sqI8g0zioTOVCK528IG+IbMdixfneTcQ7JmGRFpz7QPgmd+NZDTqBMV0gdKmYcUFN4E5Rpe43RW1NK253N/jwXD71z1rOoWaRNjK6uqusVyicyDTdUnI9LFe5Eb9RaMe1cJxrakrauIxOF0xg5zWJtveiU9DbklIerFQZqxN1cnLksMHmtJJmQwnlAd//5o9jelrMLbKGuqt1uV7amL61vq6kvVNg6fXVbzspN7DQjjuS78Gk9LZgTtbDDtiVfKlulkopVVZEXA1y/zXSgADZPJjsrb8vZlZflzWue62yIHZozpPcUxcmyeQE+W2gX8g2nuP2WTT15JtKPWDmNKi28JOkwRWkBF1FiwGkUTPZaFsCx4cJh4Dt89Hoc+fX3njzHeVnAFr3KC3axHIuDkzl/T7TnC5WSHDtti7rPJn62acbzqN+//wIF4JTCk+95WHF44OmHI7M/m/08NxkC/47FT1qGIadqHI/1+einHcew5moWdjmO479zt51r",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},k=void 0,I={},v=[];function D(e){const n={p:"p",...(0,o.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Get Container Daemon Information"}),"\n",(0,i.jsx)(s(),{method:"get",path:"/container-svc/daemon/info",context:"endpoint"}),"\n",(0,i.jsx)(n.p,{children:"Retrieve detailed information about the availability and status of container daemons on the node."}),"\n",(0,i.jsx)(c(),{parameters:void 0}),"\n",(0,i.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"Service Information",content:{"application/json":{schema:{properties:{address:{type:"string"},available:{type:"boolean"},error:{type:"string"}},required:["available"],type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function h(e={}){const{wrapper:n}={...(0,o.R)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(D,{...e})}):D(e)}}}]);