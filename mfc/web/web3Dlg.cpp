
// web3Dlg.cpp: 实现文件
/*
https://blog.csdn.net/cai747/article/details/77621336  -- set rect
https://bbs.csdn.net/topics/350250241 --屏蔽右键
*/
//

#include "stdafx.h"
#include "web3.h"
#include "web3Dlg.h"
#include "afxdialogex.h"
#include "Cui_Baseinfo.h"

#ifdef _DEBUG
#define new DEBUG_NEW
#endif

#define WM_USER_size WM_USER+100
// 用于应用程序“关于”菜单项的 CAboutDlg 对话框

class CAboutDlg : public CDialogEx
{
public:
	CAboutDlg();

// 对话框数据
#ifdef AFX_DESIGN_TIME
	enum { IDD = IDD_ABOUTBOX };
#endif

	protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV 支持

// 实现
protected:
	DECLARE_MESSAGE_MAP()
};

CAboutDlg::CAboutDlg() : CDialogEx(IDD_ABOUTBOX)
{
}

void CAboutDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
}

BEGIN_MESSAGE_MAP(CAboutDlg, CDialogEx)
END_MESSAGE_MAP()


// Cweb3Dlg 对话框



Cweb3Dlg::Cweb3Dlg(CWnd* pParent /*=nullptr*/)
	: CDialogEx(IDD_WEB3_DIALOG, pParent)
{
	m_hIcon = AfxGetApp()->LoadIcon(IDR_MAINFRAME);
}

void Cweb3Dlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
	DDX_Control(pDX, IDC_EXPLORER2,  m_web);
}

BEGIN_MESSAGE_MAP(Cweb3Dlg, CDialogEx)
	ON_WM_SYSCOMMAND()
	ON_WM_PAINT()
	ON_WM_QUERYDRAGICON()
	ON_WM_SIZE()
END_MESSAGE_MAP()


// Cweb3Dlg 消息处理程序

BOOL Cweb3Dlg::OnInitDialog()
{
	CDialogEx::OnInitDialog();

	// 将“关于...”菜单项添加到系统菜单中。

	// IDM_ABOUTBOX 必须在系统命令范围内。
	ASSERT((IDM_ABOUTBOX & 0xFFF0) == IDM_ABOUTBOX);
	ASSERT(IDM_ABOUTBOX < 0xF000);

	CMenu* pSysMenu = GetSystemMenu(FALSE);
	if (pSysMenu != nullptr)
	{
		BOOL bNameValid;
		CString strAboutMenu;
		bNameValid = strAboutMenu.LoadString(IDS_ABOUTBOX);
		ASSERT(bNameValid);
		if (!strAboutMenu.IsEmpty())
		{
			pSysMenu->AppendMenu(MF_SEPARATOR);
			pSysMenu->AppendMenu(MF_STRING, IDM_ABOUTBOX, strAboutMenu);
		}
	}

	// 设置此对话框的图标。  当应用程序主窗口不是对话框时，框架将自动
	//  执行此操作
	SetIcon(m_hIcon, TRUE);			// 设置大图标
	SetIcon(m_hIcon, FALSE);		// 设置小图标
	ModifyStyle(0, WS_CLIPCHILDREN);
	// TODO: 在此添加额外的初始化代码
	POS->setBodyRect(512, 875);
	POS->fullBodyRect(&m_web);
	COleVariant varEmpty;
	//CString str("http://loan.guiruntang.club");
	CString str("http://loan.guiruntang.club");
	COleVariant varUrl(str);
	m_web.Navigate2(varUrl, varEmpty, varEmpty, varEmpty, varEmpty); 
	return TRUE;  // 除非将焦点设置到控件，否则返回 TRUE
}

void Cweb3Dlg::OnSysCommand(UINT nID, LPARAM lParam)
{
	if ((nID & 0xFFF0) == IDM_ABOUTBOX)
	{
		CAboutDlg dlgAbout;
		dlgAbout.DoModal();
	}
	else
	{
		CDialogEx::OnSysCommand(nID, lParam);
	}
}

// 如果向对话框添加最小化按钮，则需要下面的代码
//  来绘制该图标。  对于使用文档/视图模型的 MFC 应用程序，
//  这将由框架自动完成。

void Cweb3Dlg::OnPaint()
{
	if (IsIconic())
	{
		CPaintDC dc(this); // 用于绘制的设备上下文

		SendMessage(WM_ICONERASEBKGND, reinterpret_cast<WPARAM>(dc.GetSafeHdc()), 0);

		// 使图标在工作区矩形中居中
		int cxIcon = GetSystemMetrics(SM_CXICON);
		int cyIcon = GetSystemMetrics(SM_CYICON);
		CRect rect;
		GetClientRect(&rect);
		int x = (rect.Width() - cxIcon + 1) / 2;
		int y = (rect.Height() - cyIcon + 1) / 2;

		// 绘制图标
		dc.DrawIcon(x, y, m_hIcon);
	}
	else
	{
		CDialogEx::OnPaint();
	}
}

//当用户拖动最小化窗口时系统调用此函数取得光标
//显示。
HCURSOR Cweb3Dlg::OnQueryDragIcon()
{
	return static_cast<HCURSOR>(m_hIcon);
}




BOOL Cweb3Dlg::PreTranslateMessage(MSG* pMsg)
{
	// TODO: 在此添加专用代码和/或调用基类

	// TODO: 在此添加专用代码和/或调用基类
	if (WM_RBUTTONDOWN == pMsg->message || WM_LBUTTONDBLCLK == pMsg->message)
	{
		return TRUE;
	}
	if (WM_MOUSEMOVE == pMsg->message)
	{
		CPoint point = (pMsg->pt);
		ScreenToClient(&point);
		CRect rect;
		GetClientRect(&rect);
		if (point.x <= rect.left + 3)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZEWE)));
			m_nHitTest = HTLEFT;
		}
		else if (point.x >= rect.right - 3)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZEWE)));
			m_nHitTest = HTRIGHT;
		}
		else if (point.y <= rect.top + 3)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENS)));
			m_nHitTest = HTTOP;
		}
		else if (point.y >= rect.bottom - 3)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENS)));
			m_nHitTest = HTBOTTOM;
		}
		else if (point.x <= rect.left + 5 && point.y <= rect.top + 5)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENWSE)));
			m_nHitTest = HTTOPLEFT;
		}
		else if (point.x >= rect.right - 5 && point.y <= rect.top + 5)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENESW)));
			m_nHitTest = HTTOPRIGHT;
		}
		else if (point.x <= rect.left + 5 && point.y >= rect.bottom - 5)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENESW)));
			m_nHitTest = HTBOTTOMLEFT;
		}
		else if (point.x >= rect.right - 5 && point.y >= rect.bottom - 5)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENWSE)));
			m_nHitTest = HTBOTTOMRIGHT;
		}
		else
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_ARROW)));
			m_nHitTest = 0;
		}
	} 
	if (WM_LBUTTONDOWN == pMsg->message)
	{
		if (m_nHitTest == 0) {
			return CDialogEx::PreTranslateMessage(pMsg);
		}
		CPoint point = pMsg->pt;
		if (m_nHitTest == HTTOP)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENS)));  
			 SendMessage(WM_SYSCOMMAND, SC_SIZE | WMSZ_TOP, MAKELPARAM(point.x, point.y)); 
		}
		else if (m_nHitTest == HTBOTTOM)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENS))); 
			SendMessage(WM_SYSCOMMAND, SC_SIZE | WMSZ_BOTTOM, MAKELPARAM(point.x, point.y)); 
		}
		else if (m_nHitTest == HTLEFT)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZEWE))); 
			SendMessage(WM_SYSCOMMAND, SC_SIZE | WMSZ_LEFT, MAKELPARAM(point.x, point.y)); 
		}
		else if (m_nHitTest == HTRIGHT)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZEWE))); 
			SendMessage(WM_SYSCOMMAND, SC_SIZE | WMSZ_RIGHT, MAKELPARAM(point.x, point.y)); 
		}
		else if (m_nHitTest == HTTOPLEFT)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENWSE))); 
			SendMessage(WM_SYSCOMMAND, SC_SIZE | WMSZ_RIGHT, MAKELPARAM(point.x, point.y)); 
		}
		else if (m_nHitTest == HTTOPRIGHT)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENESW))); 
			SendMessage(WM_SYSCOMMAND, SC_SIZE | WMSZ_RIGHT, MAKELPARAM(point.x, point.y)); 
		}
		else if (m_nHitTest == HTBOTTOMLEFT)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENESW))); 
			SendMessage(WM_SYSCOMMAND, SC_SIZE | WMSZ_BOTTOMLEFT, MAKELPARAM(point.x, point.y)); 
		}
		else if (m_nHitTest == HTBOTTOMRIGHT)
		{
			SetCursor(LoadCursor(NULL, MAKEINTRESOURCE(IDC_SIZENWSE))); 
			SendMessage(WM_SYSCOMMAND, SC_SIZE | WMSZ_BOTTOMRIGHT, MAKELPARAM(point.x, point.y)); 
		}
		else
		{ 
			//实现对话框跟随鼠标移动 
			PostMessage(WM_NCLBUTTONDOWN, HTCAPTION, MAKELPARAM(pMsg->pt.x, pMsg->pt.y));   //---移动当前   
		} 
	}
	return CDialogEx::PreTranslateMessage(pMsg);
}


void Cweb3Dlg::OnSize(UINT nType, int cx, int cy)
{
	CDialogEx::OnSize(nType, cx, cy);

	// TODO: 在此处添加消息处理程序代码
	POS->fullBodyRect(&m_web);
}
