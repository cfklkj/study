#pragma once
#include <windows.h>

#define POS  CBaseInfo_ui::instance()

class CBaseInfo_ui;


class CBaseInfo_ui
{
public:
	CBaseInfo_ui();
	~CBaseInfo_ui();
	static CBaseInfo_ui *instance();

	void setBodyRect(int W, int H);
	void fullBodyRect(CWnd*pWnd); 
};

