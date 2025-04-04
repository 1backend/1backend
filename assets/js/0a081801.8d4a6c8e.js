"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[8488],{87891:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>g,contentTitle:()=>f,default:()=>G,frontMatter:()=>j,metadata:()=>a,toc:()=>b});const a=JSON.parse('{"id":"1backend/get-default-model-status","title":"Get Default Model Status","description":"Retrieves the status of the default model.","source":"@site/docs/1backend/get-default-model-status.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-default-model-status","permalink":"/docs/1backend/get-default-model-status","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-default-model-status","title":"Get Default Model Status","description":"Retrieves the status of the default model.","sidebar_label":"Get Default Model Status","hide_title":true,"hide_table_of_contents":true,"api":"eJy1VU2P2zYQ/SsEz7LkjRug0KkbNFgsskWDdfa0ayBjaiwxS5EMh7LrCvrvxVDy5zaHosjFpsnhzJs379G9rJBU0D5qZ2UpHzEGjVskERsUFCF2JNwm/apwA52JonUVmvzFvthH/N7pMAV/Tfsz2qoyrcqtxt1X4TG0mkg7m8tMOo8BuNZ9JUtZY/x9TPoH31imcjKTAck7S0iy7OW7+Zy/LnGm+AO+MIGuBHVKIdGmM2YvM6mcjWgjXwfvjVapdPGNOEcvSTXYAq98YGBRjxXHtG/3oaoCUlrGvUdZSopB21oOmQQijPSIUO3PztfOGQTLAaGzlmPftPI4HogYdF1jIOFs6u4BOqsa4WxidxNSLxXzHg9XNIkYOhS7BseoRDxvr11sxFRSgK0EWNohg+X7M5EYEsZBhRXP5RovAx6HW8ny+dj5ZZ+nplbHDG79DVVMCa53eO+X+c1bBp4sdLFxQf+N1X+Z2tUQUoH3/6aWexsxWDBiiWGLQXwMwYX/V2nIJKHqgo57WT738gNCwHDbxUaWz6uBCYGamLtRqsutYpZajI2bpC8z6YHjZXG0TjGZbJZ2Cjo4ghJwSqW6YPhOYZwC0ziK5ftf3y1uJBc9YFoy9FG158iuifmy9yheppAXKTbOGLfDSqz3AgR5UJjUE90rWgFqVASLsU16eyIM3Jp4cLW2Am3lnbaRFaU5f4NQITNtoWUCb6c5J55PsgOvP+E+iUbbjWOcPBpQaTTYguaOCQzSb6Rt3RmIwdlcufYs9+d7sey8d4GZHUlqYvRlUdysQb2yeZRrCzbjJQu39zMLUW9RtFoFx1xrhSS8gbhxoeV2jFZoCRnPod7d5wexXeTzi2pUFsVut8tr2+Uu1MV0jwqovZkt8nnexNYwhoihpT83y7HaCSztgN+BXLsihRTMk46GQ24+jI3ITLIcRvjzfJHPZ0HliwXn9Y5iC/YM6R1GMb2zYlLjQVcXRPQnQ/z0P4Jp8hH/ioU3oNMbmWjsJ1c8y2OeBPTMF2yIsYNVJtkBHN33ayB8CmYYePt7h4GtucrkFoKGNRP4vBqygyjZSq+450kqhZ5FswXTjXq8egyGc+veffzCD+FkqJOEOVl2WHD2w5Hdn+W+tsAIgT+H7AdX+n40yDAc48ejH944+m6MZnZXwzD8AyxcwZs=","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Start the Default Model","permalink":"/docs/1backend/start-default-model"},"next":{"title":"Get a Model","permalink":"/docs/1backend/get-model"}}');var o=s(74848),n=s(28453),d=s(53746),i=s.n(d),l=s(56518),r=s.n(l),c=s(99972),u=s.n(c),p=s(25342),h=s.n(p),m=(s(44215),s(82223),s(24861));const j={id:"get-default-model-status",title:"Get Default Model Status",description:"Retrieves the status of the default model.",sidebar_label:"Get Default Model Status",hide_title:!0,hide_table_of_contents:!0,api:"eJy1VU2P2zYQ/SsEz7LkjRug0KkbNFgsskWDdfa0ayBjaiwxS5EMh7LrCvrvxVDy5zaHosjFpsnhzJs379G9rJBU0D5qZ2UpHzEGjVskERsUFCF2JNwm/apwA52JonUVmvzFvthH/N7pMAV/Tfsz2qoyrcqtxt1X4TG0mkg7m8tMOo8BuNZ9JUtZY/x9TPoH31imcjKTAck7S0iy7OW7+Zy/LnGm+AO+MIGuBHVKIdGmM2YvM6mcjWgjXwfvjVapdPGNOEcvSTXYAq98YGBRjxXHtG/3oaoCUlrGvUdZSopB21oOmQQijPSIUO3PztfOGQTLAaGzlmPftPI4HogYdF1jIOFs6u4BOqsa4WxidxNSLxXzHg9XNIkYOhS7BseoRDxvr11sxFRSgK0EWNohg+X7M5EYEsZBhRXP5RovAx6HW8ny+dj5ZZ+nplbHDG79DVVMCa53eO+X+c1bBp4sdLFxQf+N1X+Z2tUQUoH3/6aWexsxWDBiiWGLQXwMwYX/V2nIJKHqgo57WT738gNCwHDbxUaWz6uBCYGamLtRqsutYpZajI2bpC8z6YHjZXG0TjGZbJZ2Cjo4ghJwSqW6YPhOYZwC0ziK5ftf3y1uJBc9YFoy9FG158iuifmy9yheppAXKTbOGLfDSqz3AgR5UJjUE90rWgFqVASLsU16eyIM3Jp4cLW2Am3lnbaRFaU5f4NQITNtoWUCb6c5J55PsgOvP+E+iUbbjWOcPBpQaTTYguaOCQzSb6Rt3RmIwdlcufYs9+d7sey8d4GZHUlqYvRlUdysQb2yeZRrCzbjJQu39zMLUW9RtFoFx1xrhSS8gbhxoeV2jFZoCRnPod7d5wexXeTzi2pUFsVut8tr2+Uu1MV0jwqovZkt8nnexNYwhoihpT83y7HaCSztgN+BXLsihRTMk46GQ24+jI3ITLIcRvjzfJHPZ0HliwXn9Y5iC/YM6R1GMb2zYlLjQVcXRPQnQ/z0P4Jp8hH/ioU3oNMbmWjsJ1c8y2OeBPTMF2yIsYNVJtkBHN33ayB8CmYYePt7h4GtucrkFoKGNRP4vBqygyjZSq+450kqhZ5FswXTjXq8egyGc+veffzCD+FkqJOEOVl2WHD2w5Hdn+W+tsAIgT+H7AdX+n40yDAc48ejH944+m6MZnZXwzD8AyxcwZs=",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,g={},b=[];function y(e){const t={code:"code",p:"p",...(0,n.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Get Default Model Status"}),"\n",(0,o.jsx)(i(),{method:"get",path:"/model-svc/default-model/status",context:"endpoint"}),"\n",(0,o.jsx)(t.p,{children:"Retrieves the status of the default model."}),"\n",(0,o.jsxs)(t.p,{children:["Requires the ",(0,o.jsx)(t.code,{children:"model-svc:model:view"})," permission."]}),"\n",(0,o.jsx)(r(),{parameters:void 0}),"\n",(0,o.jsx)(u(),{title:"Body",body:void 0}),"\n",(0,o.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"Model status retrieved successfully",content:{"application/json":{schema:{properties:{status:{properties:{address:{type:"string"},assetsReady:{type:"boolean"},running:{description:"Running triggers onModelLaunch on the frontend.\n\tRunning is true when the model is both running and answering\n\t- fully loaded.",type:"boolean"}},required:["address","assetsReady","running"],type:"object"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function G(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,o.jsx)(t,{...e,children:(0,o.jsx)(y,{...e})}):y(e)}}}]);