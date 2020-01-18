// CWebBrowser2.cpp : 由 Microsoft Visual C++ 创建的 ActiveX 控件包装器类的定义


#include "stdafx.h"
#include "CWebBrowser2.h"

/////////////////////////////////////////////////////////////////////////////
// CWebBrowser2

IMPLEMENT_DYNCREATE(CWebBrowser2, CWnd)

// CWebBrowser2 属性

// CWebBrowser2 操作


BOOL CWebBrowser2::PreTranslateMessage(MSG* pMsg)
{
	// TODO: 在此添加专用代码和/或调用基类
	if (WM_RBUTTONDOWN == pMsg->message || WM_LBUTTONDBLCLK == pMsg->message)
	{ 
		return TRUE;
	}
	if (WM_LBUTTONDOWN == pMsg->message)
	{
		//PostMessage(WM_NCLBUTTONDOWN, HTCAPTION, MAKELPARAM(pMsg->pt.x, pMsg->pt.y));   //---移动当前 
		AfxGetApp()->GetMainWnd()->PostMessageW(WM_NCLBUTTONDOWN, HTCAPTION, MAKELPARAM(pMsg->pt.x, pMsg->pt.y));  //--移动父窗体
		return TRUE;
	}
	return CWnd::PreTranslateMessage(pMsg);
}
BEGIN_MESSAGE_MAP(CWebBrowser2, CWnd)
	ON_WM_LBUTTONDOWN()
END_MESSAGE_MAP()

 