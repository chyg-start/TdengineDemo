-- 删除表
DROP TABLE tb_concentrator_record;


-- 集中器设备数据（超级表）
CREATE
STABLE IF NOT EXISTS tb_concentrator_record (
	create_time timestamp,
    online smallint COMMENT ''在线状态：1-在线 0-离线'',
	voltage_a float COMMENT ''A相电压V'',
	voltage_b float COMMENT ''B相电压V'',
	voltage_c float COMMENT ''C相电压V'',
	current_a int COMMENT ''A相电流mA'',
	current_b int COMMENT ''B相电流mA'',
	current_c int COMMENT ''C相电流mA'',
	power_a int COMMENT ''A相功率 W'',
	power_b int COMMENT ''B相功率 W'',
	power_c int COMMENT ''C相功率 W'',
	power_s int COMMENT ''总功率 W'',
	repower_a int COMMENT ''A相无功功率'',
	repower_b int COMMENT ''B相无功功率'',
	repower_c int COMMENT ''C相无功功率'',
	repower_s int COMMENT ''总无功功率'',
	pf_a int COMMENT ''A相功率因数'',
	pf_b int COMMENT ''B相功率因数'',
	pf_c int COMMENT ''C相功率因数'',
	pf_s int COMMENT ''三相功率因数'',
	kwh float COMMENT ''电表电量(kWh)'',
	ad1_vol int COMMENT ''AD1路输入电压'',
	ad2_vol int COMMENT ''AD2路输入电压'',
	oc1 int COMMENT ''光耦 0输入电平 1为高 0为低'',
	oc2 int COMMENT ''光耦 1 输入电平 1为高 0为低'',
	relay1 int COMMENT ''继电器1状态：1闭合 0断开'',
	relay2 int COMMENT ''继电器2状态：1闭合 0断开'',
	relay3 int COMMENT ''继电器3状态：1闭合 0断开'',
	relay4 int COMMENT ''继电器4状态：1闭合 0断开''
) TAGS (
	con_id varchar(12) COMMENT ''集中器设备ID''
);


-- 插入子表数据
-- 表名(不区分大小写)：''CON_'' + ''集中器设备ID''
INSERT INTO con_0000000011C2 USING tb_concentrator_record TAGS (''0000000011C2'')
VALUES (
    ''2023-03-31 15:15:00.000'', 1, 0.0, 0.0, 0.0, 0,0,0,0,0,0,0,0,0,0,0,0,0,0,0, 0.0, 0,0,0,0,0,0,0,0);





