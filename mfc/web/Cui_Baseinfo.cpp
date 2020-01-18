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
	HWND hwnd = AfxGetApp()->GetMainWnd()->GetSafeHwnd();   //获取主窗口句柄 
	CRect old;
	GetWindowRect(hwnd, old);
	old.right =  X;
	old.bottom = Y;
	SetWindowPos(hwnd, HWND_NOTOPMOST, old.left, old.top, old.right, old.bottom, SWP_NOMOVE);
}

void CBaseInfo_ui::fullBodyRect(CWnd*pWnd)
{
	if (pWnd->m_hWnd == NULL) {
		return;
	}
	HWND hwnd = AfxGetApp()->GetMainWnd()->GetSafeHwnd();   //获取主窗口句柄 
	CRect old; 
	GetWindowRect(hwnd, old);  
	old.right = old.right - old.left;
	old.bottom = old.bottom - old.top;
	old.left = old.top = 0;
	pWnd->MoveWindow(old);  
}

void CBaseInfo_ui::resizeRect(CWnd * pWnd, int x, int y)
{
	CRect old;
	pWnd->GetWindowRect(old);
	old.right += y;
	old.bottom += x;
	pWnd->MoveWindow(old);
}
