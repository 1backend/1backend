"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6842],{56291:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>I,contentTitle:()=>f,default:()=>x,frontMatter:()=>u,metadata:()=>a,toc:()=>J});const a=JSON.parse('{"id":"1backend/get-message","title":"Get Message","description":"Fetch information about a specific chat message by its ID","source":"@site/docs/1backend/get-message.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-message","permalink":"/docs/1backend/get-message","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-message","title":"Get Message","description":"Fetch information about a specific chat message by its ID","sidebar_label":"Get Message","hide_title":true,"hide_table_of_contents":true,"api":"eJylVltv2zYU/isEX7YBiuQsK7b5ZXO3JPPaLcHcoA+JMRxTxxIbimTJIzueof8+HEmO7VgtUPTFl8Nz+c7tI7cyx6iC9qSdlWN5haRKoe3ShQpYJmDhahIgokell1oJVQKJCmOEAsViIzRFMf1dJtJ5DK3NNJdjWSD91SnJRHoIUCFhiHJ8v30Rs1frnGiWeKBSJtJChXIs+1jTXCYy4MdaB8zlmEKNiYyqxArkeCtp41k5UtC2kE0zZ+XonY0Y+fz70Yi/hkPnSKBNFLFWCmNc1sZsREAKGlfIcZWzhJbYA3hvtGoTzT5EdrM9gOEDl4F0FxSfdKR4AG/hnEGwskl2aZ0aqYBAmE9oIK1ELrXBaR5Pc7nqDkSOS20xCipRsLIAIlBlhZY64a53JcSUK05YxcFQvQBCgA3/13mXE1TetI2Jxb9YzW5uzftXPzmZnHqokNqqQJ5rhgnm9iDVroW9kVt8QEVtWHyi0/Te4ROJvg3CLY8ywSIVD/IPnYh1CfRNFLX/5UEOAaIyIOTTfMB/f/LCN1donzGVoc0YL+/eT4qhALXPP9O8OmIYCn7XyoXuGqSfQbC+WJdOrIMjPML1YK9cEJPpTsCmOoqlRtM6wsrTJj2F2Bwu0T339KAq85N2NKcSlv0wtE1TuwKjc/Hn7ObvL9mZlwi7AOcDZbJQU+mC/u/LlnIowKvhDAiDBSNmGFYYxGUILnxdpCaREVUdNG1a7nuNEDBMairl+H7ONEVQMC3K35hYZyvFXaiQStfTaMufrC4z5t6zuFJZ3/Rs+0yOjeRAjLrj2DoYtsiMU2BKF2l8fn5x8aPkiDtAM8bdkc4hrJPV2HgUD73KgxRLZ4xbY87szxcDKBRgc0HuEa0A1U2XWAZXtSPL082Jibeu0Fagzb3TltId3ZcIOYY94U/6JrdF3k8weP0GN+1I8gXVUqWzBKrtC1agOeMIBuOvUduiNkDB2VS56sD37VTMau9d4Lp2RSqJ/DjLzhegHtHmbJDxth5XYTI9s0B6haLSKjiutVYYhTdAfF1yOkYrtLHl9F2869u3YnWRjo6ixXGWrdfrtLB16kKR9XYxg8Kbs4t0lJZUmY4MQxVvlrMu2h5sXENRYEi1y1qVjOukqWWp89ddIjKRPA4d/FF6kY7Ogkovfma/3kWqwB4gvUYS+xv7KPftfgG+6oXQN5IJPvMGdHsPtlXZ9iN+L3cjLvc3ZCLH+zfAPJE8zay63S4g4l0wTcPijzUG3rF5IlcQNCy4GPza0JF/53K8BBPxM7l9+0/PjN+Jo0fJIO7dVNoN1xlMzf9kIh9xc/RoaeZNshtxhtOdT5RCTweWJ7zSHNLA9eU7mUjo13O/EOws2f1g74OoXi5UB4E/m+QTJtttt25N86zfHX3S4nmLO20u0rxpmv8BzneNlg==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Delete a Message","permalink":"/docs/1backend/delete-message"},"next":{"title":"Add Thread","permalink":"/docs/1backend/add-thread"}}');var i=s(74848),n=s(28453),o=s(53746),d=s.n(o),r=s(56518),c=s.n(r),l=s(99972),h=s.n(l),m=s(25342),p=s.n(m),g=(s(44215),s(82223),s(24861));const u={id:"get-message",title:"Get Message",description:"Fetch information about a specific chat message by its ID",sidebar_label:"Get Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVltv2zYU/isEX7YBiuQsK7b5ZXO3JPPaLcHcoA+JMRxTxxIbimTJIzueof8+HEmO7VgtUPTFl8Nz+c7tI7cyx6iC9qSdlWN5haRKoe3ShQpYJmDhahIgokell1oJVQKJCmOEAsViIzRFMf1dJtJ5DK3NNJdjWSD91SnJRHoIUCFhiHJ8v30Rs1frnGiWeKBSJtJChXIs+1jTXCYy4MdaB8zlmEKNiYyqxArkeCtp41k5UtC2kE0zZ+XonY0Y+fz70Yi/hkPnSKBNFLFWCmNc1sZsREAKGlfIcZWzhJbYA3hvtGoTzT5EdrM9gOEDl4F0FxSfdKR4AG/hnEGwskl2aZ0aqYBAmE9oIK1ELrXBaR5Pc7nqDkSOS20xCipRsLIAIlBlhZY64a53JcSUK05YxcFQvQBCgA3/13mXE1TetI2Jxb9YzW5uzftXPzmZnHqokNqqQJ5rhgnm9iDVroW9kVt8QEVtWHyi0/Te4ROJvg3CLY8ywSIVD/IPnYh1CfRNFLX/5UEOAaIyIOTTfMB/f/LCN1donzGVoc0YL+/eT4qhALXPP9O8OmIYCn7XyoXuGqSfQbC+WJdOrIMjPML1YK9cEJPpTsCmOoqlRtM6wsrTJj2F2Bwu0T339KAq85N2NKcSlv0wtE1TuwKjc/Hn7ObvL9mZlwi7AOcDZbJQU+mC/u/LlnIowKvhDAiDBSNmGFYYxGUILnxdpCaREVUdNG1a7nuNEDBMairl+H7ONEVQMC3K35hYZyvFXaiQStfTaMufrC4z5t6zuFJZ3/Rs+0yOjeRAjLrj2DoYtsiMU2BKF2l8fn5x8aPkiDtAM8bdkc4hrJPV2HgUD73KgxRLZ4xbY87szxcDKBRgc0HuEa0A1U2XWAZXtSPL082Jibeu0Fagzb3TltId3ZcIOYY94U/6JrdF3k8weP0GN+1I8gXVUqWzBKrtC1agOeMIBuOvUduiNkDB2VS56sD37VTMau9d4Lp2RSqJ/DjLzhegHtHmbJDxth5XYTI9s0B6haLSKjiutVYYhTdAfF1yOkYrtLHl9F2869u3YnWRjo6ixXGWrdfrtLB16kKR9XYxg8Kbs4t0lJZUmY4MQxVvlrMu2h5sXENRYEi1y1qVjOukqWWp89ddIjKRPA4d/FF6kY7Ogkovfma/3kWqwB4gvUYS+xv7KPftfgG+6oXQN5IJPvMGdHsPtlXZ9iN+L3cjLvc3ZCLH+zfAPJE8zay63S4g4l0wTcPijzUG3rF5IlcQNCy4GPza0JF/53K8BBPxM7l9+0/PjN+Jo0fJIO7dVNoN1xlMzf9kIh9xc/RoaeZNshtxhtOdT5RCTweWJ7zSHNLA9eU7mUjo13O/EOws2f1g74OoXi5UB4E/m+QTJtttt25N86zfHX3S4nmLO20u0rxpmv8BzneNlg==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},f=void 0,I={},J=[];function Y(e){const t={p:"p",...(0,n.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Get Message"}),"\n",(0,i.jsx)(d(),{method:"get",path:"/chat-svc/message/{messageId}",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Fetch information about a specific chat message by its ID"}),"\n",(0,i.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(c(),{parameters:[{description:"Message ID",in:"path",name:"messageId",required:!0,schema:{type:"string"}}]}),"\n",(0,i.jsx)(h(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(p(),{id:void 0,label:void 0,responses:{200:{description:"Message details successfully retrieved",content:{"application/json":{schema:{properties:{exists:{type:"boolean"},message:{properties:{createdAt:{type:"string"},fileIds:{description:"FileIds defines the file attachments the message has.",items:{type:"string"},type:"array"},id:{example:"msg_emSOPlW58o",type:"string"},meta:{additionalProperties:!0,type:"object"},text:{description:'Text content of the message eg. "Hi, what\'s up?"',type:"string"},threadId:{description:"ThreadId of the message.",example:"thr_emSOeEUWAg",type:"string"},updatedAt:{type:"string"},userId:{description:"UserId is the id of the user who wrote the message.\nFor AI messages this field is empty.",type:"string"}},required:["id","threadId"],type:"object"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function x(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(Y,{...e})}):Y(e)}}}]);