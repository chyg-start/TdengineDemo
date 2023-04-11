
-- 删除表
DROP TABLE tb_node_device_record;


-- 节点设备数据（超级表）
CREATE STABLE IF NOT EXISTS tb_node_device_record (
	create_time timestamp,
    con_id varchar(12) COMMENT '归属集中器ID',
    online smallint COMMENT '在线状态：1-在线 0-离线',
    alarm_state int unsigned COMMENT '报警状态位(整数，高字节在前)',
    power_status smallint COMMENT '电源状态：1-通电 0-断电',
    brightness int COMMENT '亮度 0~100',
    temperature int COMMENT '温度',
    rated_power int COMMENT '电源额定功率',
    pf int COMMENT '功率因数',
    kwh double COMMENT '电能总电量 kwh',
    input_voltage float COMMENT '输入电压V',
    input_current float COMMENT '输入电流A',
    input_power float COMMENT '输入功率P',
    leak_current float COMMENT '漏电流值 0.1mA',
    pole_leakage_vol float COMMENT '灯杆漏电电压0.1V',
    pole_leakage_cur float COMMENT '灯杆漏电电流0.1mA'
) TAGS (
	node_id varchar(12) COMMENT '节点设备ID',
	led_no int COMMENT 'LED编号(从1开始)',
	node_type varchar(2) COMMENT '节点类型'
);


-- 插入子表数据
-- 表名(不区分大小写)：'NODE_' + '节点设备ID' + 'LED编号'
INSERT INTO NODE_00000050030F_1 USING tb_node_device_record TAGS ('00000050030F', 1, '1D') VALUES (
    '2023-03-31 15:15:00.000', '0000000012C0', 1, 512, 1, 100, 30, 0, 0, 1.001, 220, 5, 50, 20.1, 0, 0);





