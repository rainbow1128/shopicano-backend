package templates

import (
	"bytes"
	"html/template"
)

var invoiceTemplate = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
   <head>
      <!--[if gte mso 9]>
      <xml>
         <o:OfficeDocumentSettings>
            <o:AllowPNG/>
            <o:PixelsPerInch>96</o:PixelsPerInch>
         </o:OfficeDocumentSettings>
      </xml>
      <![endif]-->
      <meta http-equiv="Content-type" content="text/html; charset=utf-8"/>
      <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"/>
      <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
      <meta name="format-detection" content="date=no"/>
      <meta name="format-detection" content="address=no"/>
      <meta name="format-detection" content="telephone=no"/>
      <meta name="x-apple-disable-message-reformatting"/>
      <!--[if !mso]><!-->
      <link href="https://fonts.googleapis.com/css?family=PT+Sans:400,400i,700,700i&display=swap" rel="stylesheet"/>
      <!--<![endif]-->
      <title>Shopicano : Order Details</title>
      <!--[if gte mso 9]>
      <style type="text/css" media="all">
         sup {
         font-size: 100% !important;
         }
      </style>
      <![endif]-->
      <!-- body, html, table, thead, tbody, tr, td, div, a, span { font-family: Arial, sans-serif !important; } -->
      <style type="text/css" media="screen">
         body {
         padding: 0 !important;
         margin: 0 auto !important;
         display: block !important;
         min-width: 100% !important;
         width: 100% !important;
         background: #f4ecfa;
         -webkit-text-size-adjust: none
         }
         a {
         color: #f3189e;
         text-decoration: none
         }
         p {
         padding: 0 !important;
         margin: 0 !important
         }
         img {
         margin: 0 !important;
         -ms-interpolation-mode: bicubic; /* Allow smoother rendering of resized image in Internet Explorer */
         }
         a[x-apple-data-detectors] {
         color: inherit !important;
         text-decoration: inherit !important;
         font-size: inherit !important;
         font-family: inherit !important;
         font-weight: inherit !important;
         line-height: inherit !important;
         }
         .btn-16 a {
         display: block;
         padding: 15px 35px;
         text-decoration: none;
         }
         .btn-20 a {
         display: block;
         padding: 15px 35px;
         text-decoration: none;
         }
         .l-white a {
         color: #ffffff;
         }
         .l-black a {
         color: #282828;
         }
         .l-pink a {
         color: #f3189e;
         }
         .l-grey a {
         color: #6e6e6e;
         }
         .l-purple a {
         color: #9128df;
         }
         .gradient {
         background: linear-gradient(to right, #9028df 0%, #f3189e 100%);
         }
         .btn-secondary {
         border-radius: 10px;
         background: linear-gradient(to right, #9028df 0%, #f3189e 100%);
         }
         /* Mobile styles */
         @media only screen and (max-device-width: 480px), only screen and (max-width: 480px) {
         .mpx-10 {
         padding-left: 10px !important;
         padding-right: 10px !important;
         }
         .mpx-15 {
         padding-left: 15px !important;
         padding-right: 15px !important;
         }
         .mpb-15 {
         padding-bottom: 15px !important;
         }
         u + .body .gwfw {
         width: 100% !important;
         width: 100vw !important;
         }
         .td,
         .m-shell {
         width: 100% !important;
         min-width: 100% !important;
         }
         .mt-left {
         text-align: left !important;
         }
         .mt-center {
         text-align: center !important;
         }
         .mt-right {
         text-align: right !important;
         }
         .me-left {
         margin-right: auto !important;
         }
         .me-center {
         margin: 0 auto !important;
         }
         .me-right {
         margin-left: auto !important;
         }
         .mh-auto {
         height: auto !important;
         }
         .mw-auto {
         width: auto !important;
         }
         .fluid-img img {
         width: 100% !important;
         max-width: 100% !important;
         height: auto !important;
         }
         .column,
         .column-top,
         .column-dir-top {
         float: left !important;
         width: 100% !important;
         display: block !important;
         }
         .m-hide {
         display: none !important;
         width: 0 !important;
         height: 0 !important;
         font-size: 0 !important;
         line-height: 0 !important;
         min-height: 0 !important;
         }
         .m-block {
         display: block !important;
         }
         .mw-15 {
         width: 15px !important;
         }
         .mw-2p {
         width: 2% !important;
         }
         .mw-32p {
         width: 32% !important;
         }
         .mw-49p {
         width: 49% !important;
         }
         .mw-50p {
         width: 50% !important;
         }
         .mw-100p {
         width: 100% !important;
         }
         .mmt-0 {
         margin-top: 0 !important;
         }
         }
      </style>
   </head>
   <body class="body"
      style="padding:0 !important; margin:0 auto !important; display:block !important; min-width:100% !important; width:100% !important; background:#f4ecfa; -webkit-text-size-adjust:none;">
      <center>
         <table width="100%" border="0" cellspacing="0" cellpadding="0"
            style="margin: 0; padding: 0; width: 100%; height: 100%;" bgcolor="#f4ecfa" class="gwfw">
            <tr>
               <td style="margin: 0; padding: 0; width: 100%; height: 100%;" align="center" valign="top">
                  <table width="600" border="0" cellspacing="0" cellpadding="0" class="m-shell">
                     <tr>
                        <td class="td"
                           style="width:600px; min-width:600px; font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal;">
                           <table width="100%" border="0" cellspacing="0" cellpadding="0">
                              <tr>
                                 <td class="mpx-10">
                                    <!-- Container -->
                                    <table width="100%" border="0" cellspacing="0" cellpadding="0">
                                       <tr>
                                          <td class="gradient pt-10"
                                             style="border-radius: 10px 10px 0 0; padding-top: 10px;"
                                             bgcolor="#f3189e">
                                             <table width="100%" border="0" cellspacing="0" cellpadding="0">
                                                <tr>
                                                   <td style="border-radius: 10px 10px 0 0;" bgcolor="#ffffff">
                                                      <!-- Main -->
                                                      <table width="100%" border="0" cellspacing="0"
                                                         cellpadding="0">
                                                         <tr>
                                                            <td class="px-50 mpx-15"
                                                               style="padding-left: 50px; padding-right: 50px; padding-top: 100px;">
                                                               <!-- Section - Intro -->
                                                               <table width="100%" border="0"
                                                                  cellspacing="0" cellpadding="0">
                                                                  <tr>
                                                                     <td class="pb-50"
                                                                        style="padding-bottom: 50px;">
                                                                        <table width="100%" border="0"
                                                                           cellspacing="0"
                                                                           cellpadding="0">
                                                                           <tr>
                                                                              <td class="title-36 a-center pb-15"
                                                                                 style="font-size:36px; line-height:40px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; text-align:center; padding-bottom: 15px;">
                                                                                 <strong>{{ .greetings }}</strong><br>
                                                                                 {{ .intros }}
                                                                              </td>
                                                                           </tr>
                                                                        </table>
                                                                     </td>
                                                                  </tr>
                                                               </table>
                                                               <!-- END Section - Intro -->
                                                               <!-- Section - Separator Line -->
                                                               <table width="100%" border="0"
                                                                  cellspacing="0" cellpadding="0">
                                                                  <tr>
                                                                     <td class="pb-50"
                                                                        style="padding-bottom: 50px;">
                                                                        <table width="100%" border="0"
                                                                           cellspacing="0"
                                                                           cellpadding="0">
                                                                           <tr>
                                                                              <td class="img"
                                                                                 height="1"
                                                                                 bgcolor="#ebebeb"
                                                                                 style="font-size:0pt; line-height:0pt; text-align:left;">
                                                                                 &nbsp;
                                                                              </td>
                                                                           </tr>
                                                                        </table>
                                                                     </td>
                                                                  </tr>
                                                               </table>
                                                               <!-- END Section - Separator Line -->
                                                               <!-- Section - Order Details -->
                                                               <table width="100%" border="0"
                                                                  cellspacing="0" cellpadding="0">
                                                                  <tr>
                                                                     <td class="pb-50"
                                                                        style="padding-bottom: 50px;">
                                                                        <table width="100%" border="0"
                                                                           cellspacing="0"
                                                                           cellpadding="0">
                                                                           <tr>
                                                                              <td class="pb-30"
                                                                                 style="padding-bottom: 30px;">
                                                                                 <table width="100%"
                                                                                    border="0"
                                                                                    cellspacing="0"
                                                                                    cellpadding="0">
                                                                                    <tr>
                                                                                       <th class="column-top"
                                                                                          valign="top"
                                                                                          width="230"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;">
                                                                                          <table width="100%"
                                                                                             border="0"
                                                                                             cellspacing="0"
                                                                                             cellpadding="0">
                                                                                             <tr>
                                                                                                <td class="title-20 pb-15"
                                                                                                   style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; padding-bottom: 15px;">
                                                                                                   <strong>OrderID
                                                                                                   <span class="c-purple"
                                                                                                      style="color:#9128df;">#{{ .orderHash }}</span></strong>
                                                                                                </td>
                                                                                             </tr>
                                                                                          </table>
                                                                                       </th>
                                                                                       <th class="column-top mpb-15"
                                                                                          valign="top"
                                                                                          width="30"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;"></th>
                                                                                       <th class="column-top"
                                                                                          valign="top"
                                                                                          width="230"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;">
                                                                                          <table width="100%"
                                                                                             border="0"
                                                                                             cellspacing="0"
                                                                                             cellpadding="0">
                                                                                             <tr>
                                                                                                <td class="title-20 pb-10"
                                                                                                   style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; padding-bottom: 10px;">
                                                                                                   <strong>Billing Details</strong>
                                                                                                </td>
                                                                                             </tr>
                                                                                             <tr>
                                                                                                <td class="text-16"
                                                                                                   style="font-size:16px; line-height:20px; color:#6e6e6e; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important;">
                                                                                                   {{ .billingAddress }}
                                                                                                </td>
                                                                                             </tr>
                                                                                          </table>
                                                                                       </th>
                                                                                    </tr>
                                                                                 </table>
                                                                              </td>
                                                                           </tr>
                                                                           <tr>
                                                                              <td class="py-15"
                                                                                 style="border: 1px solid #000001; border-left: 0; border-right: 0; padding-top: 15px; padding-bottom: 15px;">
                                                                                 <table width="100%"
                                                                                    border="0"
                                                                                    cellspacing="0"
                                                                                    cellpadding="0">
                                                                                    <tr>
                                                                                       <td class="title-20 mw-auto"
                                                                                          width="200"
                                                                                          style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important;">
                                                                                          <strong>Item</strong>
                                                                                       </td>
                                                                                       <td class="img"
                                                                                          width="20"
                                                                                          style="font-size:0pt; line-height:0pt; text-align:left;"></td>
                                                                                       <td class="title-20 a-center mw-auto"
                                                                                          width="40"
                                                                                          style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; text-align:center;">
                                                                                          <strong>Price</strong>
                                                                                       </td>
                                                                                       <td class="img"
                                                                                          width="20"
                                                                                          style="font-size:0pt; line-height:0pt; text-align:left;"></td>
                                                                                       <td class="title-20 a-center mw-auto"
                                                                                          width="40"
                                                                                          style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; text-align:center;">
                                                                                          <strong>Qty</strong>
                                                                                       </td>
                                                                                       <td class="img"
                                                                                          width="20"
                                                                                          style="font-size:0pt; line-height:0pt; text-align:left;"></td>
                                                                                       <td class="title-20 a-right mw-auto"
                                                                                          style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; text-align:right;">
                                                                                          <strong>Sub Total</strong>
                                                                                       </td>
                                                                                    </tr>
                                                                                 </table>
                                                                              </td>
                                                                           </tr>
                                                                           {{ range $item := .orderedItems }}
                                                                           <tr>
                                                                              <td class="py-25"
                                                                                 style="border-bottom: 1px solid #ebebeb; padding-top: 25px; padding-bottom: 25px;">
                                                                                 <table width="100%"
                                                                                    border="0"
                                                                                    cellspacing="0"
                                                                                    cellpadding="0">
                                                                                    <tr>
                                                                                       <th class="column-top"
                                                                                          valign="top"
                                                                                          width="200"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;">
                                                                                          <table width="100%"
                                                                                             border="0"
                                                                                             cellspacing="0"
                                                                                             cellpadding="0">
                                                                                             <tr>
                                                                                                <td class="title-20 pb-10"
                                                                                                   style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; padding-bottom: 10px;">
                                                                                                   <strong>{{ $item.name }}</strong>
                                                                                                </td>
                                                                                             </tr>
                                                                                          </table>
                                                                                       </th>
                                                                                       <th class="column-top mpb-15"
                                                                                          valign="top"
                                                                                          width="20"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;"></th>
                                                                                       <th class="column-top"
                                                                                          valign="top"
                                                                                          width="40"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;">
                                                                                          <table width="100%"
                                                                                             border="0"
                                                                                             cellspacing="0"
                                                                                             cellpadding="0">
                                                                                             <tr>
                                                                                                <td class="title-20 a-center mt-left"
                                                                                                   style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; text-align:center;">
                                                                                                   <strong>{{ $item.price }}</strong>
                                                                                                </td>
                                                                                             </tr>
                                                                                          </table>
                                                                                       </th>
                                                                                       <th class="column-top mpb-15"
                                                                                          valign="top"
                                                                                          width="20"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;"></th>
                                                                                       <th class="column-top"
                                                                                          valign="top"
                                                                                          width="40"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;">
                                                                                          <table width="100%"
                                                                                             border="0"
                                                                                             cellspacing="0"
                                                                                             cellpadding="0">
                                                                                             <tr>
                                                                                                <td class="title-20 a-center mt-left"
                                                                                                   style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; text-align:center;">
                                                                                                   <strong>&times;{{ $item.quantity }}</strong>
                                                                                                </td>
                                                                                             </tr>
                                                                                          </table>
                                                                                       </th>
                                                                                       <th class="column-top mpb-15"
                                                                                          valign="top"
                                                                                          width="20"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;"></th>
                                                                                       <th class="column-top"
                                                                                          valign="top"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;">
                                                                                          <table width="100%"
                                                                                             border="0"
                                                                                             cellspacing="0"
                                                                                             cellpadding="0">
                                                                                             <tr>
                                                                                                <td class="title-20 a-right mt-left"
                                                                                                   style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; text-align:right;">
                                                                                                   <strong>{{ $item.subTotal }}</strong>
                                                                                                </td>
                                                                                             </tr>
                                                                                          </table>
                                                                                       </th>
                                                                                    </tr>
                                                                                 </table>
                                                                              </td>
                                                                           </tr>
                                                                           {{end}}
                                                                           <tr>
                                                                              <td class="pt-30"
                                                                                 style="border-top: 1px solid #000001; padding-top: 30px;">
                                                                                 <table width="100%"
                                                                                    border="0"
                                                                                    cellspacing="0"
                                                                                    cellpadding="0">
                                                                                    <tr>
                                                                                       {{ if .isShippable }}
                                                                                       <th class="column-top"
                                                                                          valign="top"
                                                                                          width="230"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;">
                                                                                          <table width="100%"
                                                                                             border="0"
                                                                                             cellspacing="0"
                                                                                             cellpadding="0">
                                                                                             <tr>
                                                                                                <td class="title-20 pb-10"
                                                                                                   style="font-size:20px; line-height:24px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; padding-bottom: 10px;">
                                                                                                   <strong>Shipping Address</strong>
                                                                                                </td>
                                                                                             </tr>
                                                                                             <tr>
                                                                                                <td class="text-16"
                                                                                                   style="font-size:16px; line-height:20px; color:#6e6e6e; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important;">
                                                                                                   {{ .shippingAddress }}
                                                                                                </td>
                                                                                             </tr>
                                                                                          </table>
                                                                                       </th>
                                                                                       <th class="column-top mpb-15"
                                                                                          valign="top"
                                                                                          width="30"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;"></th>
                                                                                       <th class="column-top"
                                                                                          valign="top"
                                                                                          width="230"
                                                                                          style="font-size:0pt; line-height:0pt; padding:0; margin:0; font-weight:normal; vertical-align:top;">
                                                                                          <table width="100%"
                                                                                             border="0"
                                                                                             cellspacing="0"
                                                                                             cellpadding="0">
                                                                                             <tr>
                                                                                                <td align="right">
                                                                                                   <table border="0"
                                                                                                      cellspacing="0"
                                                                                                      cellpadding="0"
                                                                                                      class="mw-100p">
                                                                                                      <tr>
                                                                                                         <td class="title-20 lh-30 a-right mt-left mw-auto"
                                                                                                            width="100"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; line-height: 30px; text-align:right;">
                                                                                                            <strong>Subtotal:</strong>
                                                                                                         </td>
                                                                                                         <td class="img mw-15"
                                                                                                            width="20"
                                                                                                            style="font-size:0pt; line-height:0pt; text-align:left;"></td>
                                                                                                         <td class="title-20 lh-30 mt-right"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; line-height: 30px;">
                                                                                                            {{ .subTotal }}
                                                                                                         </td>
                                                                                                      </tr>
                                                                                                      {{ if .isShippable }}
                                                                                                      <tr>
                                                                                                         <td class="title-20 lh-30 a-right mt-left"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; line-height: 30px; text-align:right;">
                                                                                                            <strong>SC:</strong>
                                                                                                         </td>
                                                                                                         <td class="img mw-15"
                                                                                                            style="font-size:0pt; line-height:0pt; text-align:left;"></td>
                                                                                                         <td class="title-20 lh-30 mt-right"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; line-height: 30px;">
                                                                                                            {{ .shippingCharge }}
                                                                                                         </td>
                                                                                                      </tr>
                                                                                                      <tr>
                                                                                                         <td class="title-20 lh-30 a-right mt-left"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; line-height: 30px; text-align:right;">
                                                                                                            <strong>PPF:</strong>
                                                                                                         </td>
                                                                                                         <td class="img mw-15"
                                                                                                            style="font-size:0pt; line-height:0pt; text-align:left;"></td>
                                                                                                         <td class="title-20 lh-30 mt-right"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; line-height: 30px;">
                                                                                                            {{ .paymentProcessingFee }}
                                                                                                         </td>
                                                                                                      </tr>
                                                                                                      <tr>
                                                                                                         <td class="title-20 lh-30 a-right mt-left"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; line-height: 30px; text-align:right;">
                                                                                                            <strong>PG:</strong>
                                                                                                         </td>
                                                                                                         <td class="img mw-15"
                                                                                                            style="font-size:0pt; line-height:0pt; text-align:left;"></td>
                                                                                                         <td class="title-20 lh-30 mt-right"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; line-height: 30px;">
                                                                                                            {{ .paymentGateway }}
                                                                                                         </td>
                                                                                                      </tr>
                                                                                                      <tr>
                                                                                                         <td class="title-20 lh-30 a-right mt-left pt-10"
                                                                                                            style="font-size:20px; color:#282828; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; line-height: 30px; text-align:right; padding-top: 10px;">
                                                                                                            <strong>TOTAL:</strong>
                                                                                                         </td>
                                                                                                         <td class="img mw-15 pt-10"
                                                                                                            style="font-size:0pt; line-height:0pt; text-align:left; padding-top: 10px;"></td>
                                                                                                         <td class="title-20 lh-30 c-purple pt-10 mt-right"
                                                                                                            style="font-size:20px; font-family:'PT Sans', Arial, sans-serif; text-align:left; min-width:auto !important; line-height: 30px; color:#9128df; padding-top: 10px;">
                                                                                                            <strong>{{ .grandTotal }}</strong>
                                                                                                         </td>
                                                                                                      </tr>
                                                                                                   </table>
                                                                                                </td>
                                                                                             </tr>
                                                                                          </table>
                                                                                       </th>
                                                                                    </tr>
                                                                                 </table>
                                                                              </td>
                                                                           </tr>
                                                                        </table>
                                                                     </td>
                                                                  </tr>
                                                               </table>
                                                               <!-- END Section - Order Details -->
                                                            </td>
                                                         </tr>
                                                      </table>
                                                      <!-- END Main -->
                                                   </td>
                                                </tr>
                                             </table>
                                          </td>
                                       </tr>
                                    </table>
                                    <!-- END Container -->
                                    <!-- Footer -->
                                    <table width="100%" border="0" cellspacing="0" cellpadding="0">
                                       <tr>
                                          <td class="p-50 mpx-15" bgcolor="#949196"
                                             style="border-radius: 0 0 10px 10px; padding: 50px;">
                                             <table width="100%" border="0" cellspacing="0" cellpadding="0">
                                                <tr>
                                                   <td class="text-14 lh-24 a-center c-white l-white pb-20"
                                                      style="font-size:14px; font-family:'PT Sans', Arial, sans-serif; min-width:auto !important; line-height: 24px; text-align:center; color:#ffffff; padding-bottom: 20px;">
                                                      <p style="font-size: 8px;">*SC = Shipping Charge | PPF = Payment Processing Fee | PG = Payment Gateway</p>
                                                      
                                                      Address name St. 12, City Name, State, Country Name
                                                      <br/>
                                                      <a target="_blank"
                                                         class="link c-white"
                                                         style="text-decoration:none; color:#ffffff;"><span
                                                         class="link c-white"
                                                         style="text-decoration:none; color:#ffffff;">(880) 171-0330000</span></a>
                                                      <br/>
                                                      <a href="mailto:hello@shopicano.com" target="_blank"
                                                         class="link c-white"
                                                         style="text-decoration:none; color:#ffffff;"><span
                                                         class="link c-white"
                                                         style="text-decoration:none; color:#ffffff;">hello@shopicano.com</span></a>
                                                      - <a href="https://www.shopicano.com" target="_blank"
                                                         class="link c-white"
                                                         style="text-decoration:none; color:#ffffff;"><span
                                                         class="link c-white"
                                                         style="text-decoration:none; color:#ffffff;">www.shopicano.com</span></a>
														<br>
														Copyright &copy 2020 Shopicano | All rights reserved.
                                                   </td>
                                                </tr>
                                             </table>
                                          </td>
                                       </tr>
                                    </table>
                                    <!-- END Footer -->
                                 </td>
                              </tr>
                           </table>
                        </td>
                     </tr>
                  </table>
               </td>
            </tr>
         </table>
      </center>
   </body>
</html>
`

func GenerateInvoiceEmailHTML(params map[string]interface{}) (string, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("InvoiceTemplate").Parse(invoiceTemplate))
	if err := t.Execute(&buf, params); err != nil {
		return "", err
	}
	return buf.String(), nil
}
