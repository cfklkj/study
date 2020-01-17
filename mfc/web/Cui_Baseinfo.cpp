#include "stdafx.h"
#include "Cui_baseinfo.h"


static CBaseInfo_ui m_ui;

CBaseInfo_ui* CBaseInfo_ui::instance() {
	return &m_ui;
}

CBaseInfo_ui::CBaseInfo_ui()
{
}


CBaseInfo_ui::~CBaseInfo_ui()
{
}
void CBaseInfo_ui::setBodyRect(int X, int Y)
{
	HWND hwnd = AfxGetApp()->GetMainWnd()->GetSafeHwnd();   //��ȡ�����ھ�� 
	CRect old;
	GetWindowRect(hwnd, old);
	old.right =  X;
	old.bottom = Y;
	SetWindowPos(hwnd, HWND_NOTOPMOST, old.left, old.top, old.right, old.bottom, SWP_NOMOVE);
}

void CBaseInfo_ui::fullBodyRect(CWnd*pWnd)
{
	HWND hwnd = AfxGetApp()->GetMainWnd()->GetSafeHwnd();   //��ȡ�����ھ�� 
	CRect old; 
	GetWindowRect(hwnd, old); 
	pWnd->MoveWindow(old, true); 
}