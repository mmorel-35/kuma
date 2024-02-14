import{d as b,l as C,a as _,o as r,c as u,m as l,e as s,w as t,q as a,t as n,f as e,F as O,E as x,X as P,b as y,p as f,W as m,_ as I,A as V,R as z,x as B,y as $,a2 as K}from"./index-sL-Jj2ZJ.js";import{a as N,K as L}from"./index-FZCiQto1.js";import{S as A}from"./StatusBadge-ft_D07JT.js";import{T as D}from"./TagList-94RPLvP4.js";import{T as k}from"./TextWithCopyButton-lp9UdEeo.js";import"./CopyButton-6B7rfdOI.js";const E={class:"stack"},U={class:"stack-with-borders"},F={class:"status-with-reason"},W={key:0},q={class:"mt-4"},G={class:"stack-with-borders"},X={class:"mt-4 stack"},Z={class:"mt-2 stack-with-borders"},j=b({__name:"DataPlaneSummary",props:{dataplaneOverview:{}},setup(p){const{t:o,formatIsoDate:w}=C(),i=p;return(R,S)=>{const g=_("KTooltip"),v=_("DataCollection"),h=_("KBadge");return r(),u("div",E,[l("div",U,[s(m,{layout:"horizontal"},{title:t(()=>[e(n(a(o)("http.api.property.status")),1)]),body:t(()=>[l("div",F,[s(A,{status:i.dataplaneOverview.status},null,8,["status"]),e(),i.dataplaneOverview.dataplane.networking.type==="standard"?(r(),y(v,{key:0,items:i.dataplaneOverview.dataplane.networking.inbounds,predicate:c=>!c.health.ready,empty:!1},{default:t(({items:c})=>[s(g,{class:"reason-tooltip",placement:"bottomEnd"},{content:t(()=>[l("ul",null,[(r(!0),u(O,null,x(c,d=>(r(),u("li",{key:`${d.service}:${d.port}`},n(a(o)("data-planes.routes.item.unhealthy_inbound",{service:d.service,port:d.port})),1))),128))])]),default:t(()=>[s(a(P),{color:a(N),size:a(L),"hide-title":""},null,8,["color","size"]),e()]),_:2},1024)]),_:1},8,["items","predicate"])):f("",!0)])]),_:1}),e(),s(m,{layout:"horizontal"},{title:t(()=>[e(n(a(o)("data-planes.routes.item.last_updated")),1)]),body:t(()=>[e(n(a(w)(i.dataplaneOverview.modificationTime)),1)]),_:1})]),e(),i.dataplaneOverview.dataplane.networking.gateway?(r(),u("div",W,[l("h3",null,n(a(o)("data-planes.routes.item.gateway")),1),e(),l("div",q,[l("div",G,[s(m,{layout:"horizontal"},{title:t(()=>[e(n(a(o)("http.api.property.tags")),1)]),body:t(()=>[s(D,{alignment:"right",tags:i.dataplaneOverview.dataplane.networking.gateway.tags},null,8,["tags"])]),_:1}),e(),s(m,{layout:"horizontal"},{title:t(()=>[e(n(a(o)("http.api.property.address")),1)]),body:t(()=>[s(k,{text:`${i.dataplaneOverview.dataplane.networking.address}`},null,8,["text"])]),_:1})])])])):f("",!0),e(),i.dataplaneOverview.dataplane.networking.type==="standard"?(r(),y(v,{key:1,items:i.dataplaneOverview.dataplane.networking.inbounds},{default:t(({items:c})=>[l("div",null,[l("h3",null,n(a(o)("data-planes.routes.item.inbounds")),1),e(),l("div",X,[(r(!0),u(O,null,x(c,(d,T)=>(r(),u("div",{key:T,class:"inbound"},[l("h4",null,[s(k,{text:d.tags["kuma.io/service"]},{default:t(()=>[e(n(a(o)("data-planes.routes.item.inbound_name",{service:d.tags["kuma.io/service"]})),1)]),_:2},1032,["text"])]),e(),l("div",Z,[s(m,{layout:"horizontal"},{title:t(()=>[e(n(a(o)("http.api.property.status")),1)]),body:t(()=>[d.health.ready?(r(),y(h,{key:0,appearance:"success"},{default:t(()=>[e(n(a(o)("data-planes.routes.item.health.ready")),1)]),_:1})):(r(),y(h,{key:1,appearance:"danger"},{default:t(()=>[e(n(a(o)("data-planes.routes.item.health.not_ready")),1)]),_:1}))]),_:2},1024),e(),s(m,{layout:"horizontal"},{title:t(()=>[e(n(a(o)("http.api.property.tags")),1)]),body:t(()=>[s(D,{alignment:"right",tags:d.tags},null,8,["tags"])]),_:2},1024),e(),s(m,{layout:"horizontal"},{title:t(()=>[e(n(a(o)("http.api.property.address")),1)]),body:t(()=>[s(k,{text:d.addressPort},null,8,["text"])]),_:2},1024)])]))),128))])])]),_:1},8,["items"])):f("",!0)])}}}),H=I(j,[["__scopeId","data-v-56c9aa06"]]),J=p=>(B("data-v-21ad478f"),p=p(),$(),p),M={class:"summary-title-wrapper"},Q=J(()=>l("img",{"aria-hidden":"true",src:K},null,-1)),Y={class:"summary-title"},tt={key:1,class:"stack"},et=b({__name:"DataPlaneSummaryView",props:{name:{},dataplaneOverview:{default:void 0}},setup(p){const{t:o}=C(),w=V(),i=p;return(R,S)=>{const g=_("RouteTitle"),v=_("RouterLink"),h=_("AppView"),c=_("RouteView");return r(),y(c,{name:a(w).name},{default:t(()=>[s(h,null,{title:t(()=>[l("div",M,[Q,e(),l("h2",Y,[s(v,{to:{name:"data-plane-detail-view",params:{dataPlane:i.name}}},{default:t(()=>[s(g,{title:a(o)("data-planes.routes.item.title",{name:i.name})},null,8,["title"])]),_:1},8,["to"])])])]),default:t(()=>[e(),i.dataplaneOverview===void 0?(r(),y(z,{key:0},{message:t(()=>[l("p",null,n(a(o)("common.collection.summary.empty_message",{type:"Data Plane Proxy"})),1)]),default:t(()=>[e(n(a(o)("common.collection.summary.empty_title",{type:"Data Plane Proxy"}))+" ",1)]),_:1})):(r(),u("div",tt,[s(H,{class:"mt-4","dataplane-overview":i.dataplaneOverview},null,8,["dataplane-overview"])]))]),_:1})]),_:1},8,["name"])}}}),rt=I(et,[["__scopeId","data-v-21ad478f"]]);export{rt as default};
