package test

var flex1 = []byte(`{
 "type": "bubble",
 "hero": {
   "type": "image",
   "url": "https://dynamic-cdn.eggdigital.com/dwFsw0P8M.png",
   "size": "full",
   "aspectRatio": "20:13",
   "aspectMode": "cover",
   "action": {
     "type": "uri",
     "uri": "https://trueonline.truecorp.co.th/package"
   }
 },
 "body": {
   "type": "box",
   "layout": "vertical",
   "contents": [
     {
       "type": "text",
       "text": "แพ็คเกจ 699/เดือน",
       "weight": "bold",
       "size": "xl"
     },
     {
       "type": "box",
       "layout": "baseline",
       "margin": "md",
       "contents": [
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
         },
         {
           "type": "text",
           "text": "4.0",
           "size": "sm",
           "color": "#999999",
           "margin": "md",
           "flex": 0
         }
       ]
     },
     {
       "type": "box",
       "layout": "vertical",
       "margin": "lg",
       "spacing": "sm",
       "contents": [
         {
           "type": "box",
           "layout": "baseline",
           "spacing": "sm",
           "contents": [
             {
               "type": "text",
               "text": "-",
               "color": "#aaaaaa",
               "size": "sm",
               "flex": 1
             },
             {
               "type": "text",
               "text": "download up to 1000 mbps",
               "wrap": true,
               "color": "#666666",
               "size": "sm",
               "flex": 5
             }
           ]
         },
         {
           "type": "box",
           "layout": "baseline",
           "spacing": "sm",
           "contents": [
             {
               "type": "text",
               "text": "-",
               "color": "#aaaaaa",
               "size": "sm",
               "flex": 1
             },
             {
               "type": "text",
               "text": "upload up to 200mbps",
               "wrap": true,
               "color": "#666666",
               "size": "sm",
               "flex": 5
             }
           ]
         }
       ]
     }
   ]
 },
 "footer": {
   "type": "box",
   "layout": "vertical",
   "spacing": "sm",
   "contents": [
     {
       "type": "button",
       "style": "link",
       "height": "sm",
       "action": {
         "type": "uri",
         "label": "สมัครเลย",
         "uri": "https://trueonline.truecorp.co.th/package"
       }
     },
     {
       "type": "button",
       "style": "link",
       "height": "sm",
       "action": {
         "type": "message",
         "label": "สอบถามเพิ่มเติม",
         "text": "ติดต่อพนักงาน"
       }
     },
     {
       "type": "spacer",
       "size": "sm"
     }
   ],
   "flex": 0
 }
}`)


var flexs = []byte(`{
 "type": "carousel",
 "contents": [
   {
 "type": "bubble",
 "hero": {
   "type": "image",
   "url": "https://dynamic-cdn.eggdigital.com/dwFsw0P8M.png",
   "size": "full",
   "aspectRatio": "20:13",
   "aspectMode": "cover",
   "action": {
     "type": "uri",
     "uri": "https://trueonline.truecorp.co.th/package"
   }
 },
 "body": {
   "type": "box",
   "layout": "vertical",
   "contents": [
     {
       "type": "text",
       "text": "แพ็คเกจ 399/เดือน",
       "weight": "bold",
       "size": "xl"
     },
     {
       "type": "box",
       "layout": "baseline",
       "margin": "md",
       "contents": [
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
         },
         {
           "type": "text",
           "text": "4.0",
           "size": "sm",
           "color": "#999999",
           "margin": "md",
           "flex": 0
         }
       ]
     },
     {
       "type": "box",
       "layout": "vertical",
       "margin": "lg",
       "spacing": "sm",
       "contents": [
         {
           "type": "box",
           "layout": "baseline",
           "spacing": "sm",
           "contents": [
             {
               "type": "text",
               "text": "-",
               "color": "#aaaaaa",
               "size": "sm",
               "flex": 1
             },
             {
               "type": "text",
               "text": "download up to 1000 mbps",
               "wrap": true,
               "color": "#666666",
               "size": "sm",
               "flex": 5
             }
           ]
         },
         {
           "type": "box",
           "layout": "baseline",
           "spacing": "sm",
           "contents": [
             {
               "type": "text",
               "text": "-",
               "color": "#aaaaaa",
               "size": "sm",
               "flex": 1
             },
             {
               "type": "text",
               "text": "upload up to 200mbps",
               "wrap": true,
               "color": "#666666",
               "size": "sm",
               "flex": 5
             }
           ]
         }
       ]
     }
   ]
 },
 "footer": {
   "type": "box",
   "layout": "vertical",
   "spacing": "sm",
   "contents": [
     {
       "type": "button",
       "style": "link",
       "height": "sm",
       "action": {
         "type": "uri",
         "label": "สมัครเลย",
         "uri": "https://trueonline.truecorp.co.th/package"
       }
     },
     {
       "type": "button",
       "style": "link",
       "height": "sm",
       "action": {
         "type": "message",
         "label": "สอบถามเพิ่มเติม",
         "text": "ติดต่อพนักงาน"
       }
     },
     {
       "type": "spacer",
       "size": "sm"
     }
   ],
   "flex": 0
 }
},
   {
 "type": "bubble",
 "hero": {
   "type": "image",
   "url": "https://dynamic-cdn.eggdigital.com/dwFsw0P8M.png",
   "size": "full",
   "aspectRatio": "20:13",
   "aspectMode": "cover",
   "action": {
     "type": "uri",
     "uri": "https://trueonline.truecorp.co.th/package"
   }
 },
 "body": {
   "type": "box",
   "layout": "vertical",
   "contents": [
     {
       "type": "text",
       "text": "แพ็คเกจ 699/เดือน",
       "weight": "bold",
       "size": "xl"
     },
     {
       "type": "box",
       "layout": "baseline",
       "margin": "md",
       "contents": [
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
         },
         {
           "type": "text",
           "text": "4.0",
           "size": "sm",
           "color": "#999999",
           "margin": "md",
           "flex": 0
         }
       ]
     },
     {
       "type": "box",
       "layout": "vertical",
       "margin": "lg",
       "spacing": "sm",
       "contents": [
         {
           "type": "box",
           "layout": "baseline",
           "spacing": "sm",
           "contents": [
             {
               "type": "text",
               "text": "-",
               "color": "#aaaaaa",
               "size": "sm",
               "flex": 1
             },
             {
               "type": "text",
               "text": "download up to 1000 mbps",
               "wrap": true,
               "color": "#666666",
               "size": "sm",
               "flex": 5
             }
           ]
         },
         {
           "type": "box",
           "layout": "baseline",
           "spacing": "sm",
           "contents": [
             {
               "type": "text",
               "text": "-",
               "color": "#aaaaaa",
               "size": "sm",
               "flex": 1
             },
             {
               "type": "text",
               "text": "upload up to 200mbps",
               "wrap": true,
               "color": "#666666",
               "size": "sm",
               "flex": 5
             }
           ]
         }
       ]
     }
   ]
 },
 "footer": {
   "type": "box",
   "layout": "vertical",
   "spacing": "sm",
   "contents": [
     {
       "type": "button",
       "style": "link",
       "height": "sm",
       "action": {
         "type": "uri",
         "label": "สมัครเลย",
         "uri": "https://trueonline.truecorp.co.th/package"
       }
     },
     {
       "type": "button",
       "style": "link",
       "height": "sm",
       "action": {
         "type": "message",
         "label": "สอบถามเพิ่มเติม",
         "text": "ติดต่อพนักงาน"
       }
     },
     {
       "type": "spacer",
       "size": "sm"
     }
   ],
   "flex": 0
 }
},
   {
 "type": "bubble",
 "hero": {
   "type": "image",
   "url": "https://dynamic-cdn.eggdigital.com/dwFsw0P8M.png",
   "size": "full",
   "aspectRatio": "20:13",
   "aspectMode": "cover",
   "action": {
     "type": "uri",
     "uri": "https://trueonline.truecorp.co.th/package"
   }
 },
 "body": {
   "type": "box",
   "layout": "vertical",
   "contents": [
     {
       "type": "text",
       "text": "แพ็คเกจ 999/เดือน",
       "weight": "bold",
       "size": "xl"
     },
     {
       "type": "box",
       "layout": "baseline",
       "margin": "md",
       "contents": [
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
         },
         {
           "type": "icon",
           "size": "sm",
           "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
         },
         {
           "type": "text",
           "text": "4.0",
           "size": "sm",
           "color": "#999999",
           "margin": "md",
           "flex": 0
         }
       ]
     },
     {
       "type": "box",
       "layout": "vertical",
       "margin": "lg",
       "spacing": "sm",
       "contents": [
         {
           "type": "box",
           "layout": "baseline",
           "spacing": "sm",
           "contents": [
             {
               "type": "text",
               "text": "-",
               "color": "#aaaaaa",
               "size": "sm",
               "flex": 1
             },
             {
               "type": "text",
               "text": "download up to 1000 mbps",
               "wrap": true,
               "color": "#666666",
               "size": "sm",
               "flex": 5
             }
           ]
         },
         {
           "type": "box",
           "layout": "baseline",
           "spacing": "sm",
           "contents": [
             {
               "type": "text",
               "text": "-",
               "color": "#aaaaaa",
               "size": "sm",
               "flex": 1
             },
             {
               "type": "text",
               "text": "upload up to 200mbps",
               "wrap": true,
               "color": "#666666",
               "size": "sm",
               "flex": 5
             }
           ]
         }
       ]
     }
   ]
 },
 "footer": {
   "type": "box",
   "layout": "vertical",
   "spacing": "sm",
   "contents": [
     {
       "type": "button",
       "style": "link",
       "height": "sm",
       "action": {
         "type": "uri",
         "label": "สมัครเลย",
         "uri": "https://trueonline.truecorp.co.th/package"
       }
     },
     {
       "type": "button",
       "style": "link",
       "height": "sm",
       "action": {
         "type": "message",
         "label": "สอบถามเพิ่มเติม",
         "text": "ติดต่อพนักงาน"
       }
     },
     {
       "type": "spacer",
       "size": "sm"
     }
   ],
   "flex": 0
 }
}
 ]
}`)

