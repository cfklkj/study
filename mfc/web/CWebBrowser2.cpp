// CWebBrowser2.cpp : �� Microsoft Visual C++ ������ ActiveX �ؼ���װ����Ķ���


#include "stdafx.h"
#include "CWebBrowser2.h"

/////////////////////////////////////////////////////////////////////////////
// CWebBrowser2

IMPLEMENT_DYNCREATE(CWebBrowser2, CWnd)

// CWebBrowser2 ����

// CWebBrowser2 ����


BOOL CWebBrowser2::PreTranslateMessage(MSG* pMsg)
{
	// TODO: �ڴ����ר�ô����/����û���
	if (WM_RBUTTONDOWN == pMsg->message || WM_LBUTTONDBLCLK == pMsg->message)
	{ 
		return TRUE;
	}
	if (WM_LBUTTONDOWN == pMsg->message)
	{
		//PostMessage(WM_NCLBUTTONDOWN, HTCAPTION, MAKELPARAM(pMsg->pt.x, pMsg->pt.y));   //---�ƶ���ǰ 
		AfxGetApp()->GetMainWnd()->PostMessageW(WM_NCLBUTTONDOWN, HTCAPTION, MAKELPARAM(pMsg->pt.x, pMsg->pt.y));  //--�ƶ�������
		return TRUE;
	}
	return CWnd::PreTranslateMessage(pMsg);
}
BEGIN_MESSAGE_MAP(CWebBrowser2, CWnd)
	ON_WM_LBUTTONDOWN()
END_MESSAGE_MAP()

 